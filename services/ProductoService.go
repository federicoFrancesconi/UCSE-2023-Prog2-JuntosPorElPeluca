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
	DescontarStockProducto(idProducto int, cantidadDescontada int) error
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

func (service *ProductoService) DescontarStockProducto(idProducto int, cantidadDescontada int) error {
	//Buscamos el producto del que hay que descontar la cantidad
	producto, err := service.repository.ObtenerProductoPorCodigo(idProducto)

	if err != nil {
		return err
	}

	//Modificamos el stock
	producto.StockActual = producto.StockActual - cantidadDescontada

	//Actualizamos la base de datos
	return service.repository.ActualizarProducto(*producto)
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto) error {
	return service.repository.EliminarProducto(producto.GetModel().CodigoProducto)
}
