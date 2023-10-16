package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type ProductoService struct {
	repository repositories.ProductoRepositoryInterface
}

type ProductoServiceInterface interface {
	CrearProducto(producto *dto.Producto) error
	//ObtenerProductos() ([]dto.Producto, error)
	ObtenerProductosFiltrados(tipoProducto model.TipoProducto) ([]dto.Producto, error)
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

func (service *ProductoService) ObtenerProductosFiltrados(tipoProducto model.TipoProducto) ([]dto.Producto, error) {
	productos, err := service.repository.ObtenerProductosFiltrados(tipoProducto)

	if err != nil {
		return nil, err
	}

	var productosDTO []dto.Producto

	for _, producto := range productos {
		productosDTO = append(productosDTO, *dto.NewProducto(producto))
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
