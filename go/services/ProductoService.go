package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
)

type ProductoService struct {
	repository repositories.ProductoRepositoryInterface
}

type ProductoServiceInterface interface {
	CrearProducto(*dto.Producto) error
	ObtenerProductosFiltrados(utils.FiltroProducto) ([]dto.Producto, error)
	ActualizarProducto(*dto.Producto) error
	EliminarProducto(*dto.Producto) error
}

func NewProductoService(repository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		repository: repository,
	}
}

func (service *ProductoService) CrearProducto(producto *dto.Producto) error {
	return service.repository.CrearProducto(producto.GetModel())
}

func (service *ProductoService) ObtenerProductosFiltrados(filtro utils.FiltroProducto) ([]dto.Producto, error) {
	productos, err := service.repository.ObtenerProductosFiltrados(filtro)

	if err != nil {
		return nil, err
	}

	var productosDTO []dto.Producto

	for _, producto := range productos {
		productosDTO = append(productosDTO, *dto.NewProducto(producto))
	}

	return productosDTO, nil
}

func (service *ProductoService) ActualizarProducto(producto *dto.Producto) error {
	return service.repository.ActualizarProducto(producto.GetModel())
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto) error {
	return service.repository.EliminarProducto(producto.GetModel())
}
