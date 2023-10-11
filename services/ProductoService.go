package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type ProductoService struct {
	repository repositories.ProductoRepositoryInterface
}

type ProductoServiceInterface interface {
	// - ActualizarStockProductos????? (para tirarle directamente la lista? nose)
	CrearProducto(producto *dto.Producto) error
	ObtenerProductos() ([]dto.Producto, error)
	ActualizarStockProducto(producto *dto.Producto) error
	EliminarProducto(producto *dto.Producto) error
}

func NewProductoService(repository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		repository: repository,
	}
}
