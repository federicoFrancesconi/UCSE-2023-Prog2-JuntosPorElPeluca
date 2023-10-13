package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type ProductoService struct {
	repository repositories.ProductoRepositoryInterface
}

type ProductoServiceInterface interface {
	CrearProducto(producto *dto.Producto) error
	ObtenerProductos() ([]dto.Producto, error)
	ActualizarStockProducto(producto *dto.Producto) error
	ActualizarStockProductos(productos *[]dto.Producto) error
	EliminarProducto(producto *dto.Producto) error
}

func NewProductoService(repository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		repository: repository,
	}
}

func (service *ProductoService) CrearProducto(producto *dto.Producto) error {
	return service.repository.CrearProducto(producto.GetModel())
}

func (service *ProductoService) ObtenerProductos() ([]dto.Producto, error) {
	productos, err := service.repository.ObtenerProductos()
	if err != nil {
		return nil, err
	}

	var productosDTO []dto.Producto

	for _, producto := range productos {
		productoDTO := *dto.NewProducto(producto)
		productosDTO = append(productosDTO, productoDTO)
	}

	return productosDTO, nil
}

func (service *ProductoService) ActualizarStockProducto(producto *dto.Producto) error {
	return service.repository.ActualizarProducto(producto.GetModel())
}

func (service *ProductoService) ActualizarStockProductos(productos *[]dto.Producto) error {
	for _, producto := range *productos {
		err := service.repository.ActualizarProducto(producto.GetModel())
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto) error {
	return service.repository.EliminarProducto(producto.GetModel().CodigoProducto)
}
