package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"errors"
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
	//Valido el tipo de producto
	if !model.EsUnTipoProductoValido(producto.TipoDeProducto) {
		return errors.New("el tipo de producto ingresado no es válido")
	}

	return service.repository.CrearProducto(producto.GetModel())
}

func (service *ProductoService) ObtenerProductosFiltrados(filtro utils.FiltroProducto) ([]dto.Producto, error) {
	//Valido el tipo de producto que usa para filtrar
	if !model.EsUnTipoProductoValido(filtro.TipoProducto) && filtro.TipoProducto != "" {
		return nil, errors.New("el tipo de producto ingresado no es válido")
	}

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
	//Valido el tipo de producto
	if !model.EsUnTipoProductoValido(producto.TipoDeProducto) {
		return errors.New("el tipo de producto ingresado no es válido")
	}

	return service.repository.ActualizarProducto(producto.GetModel())
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto) error {
	return service.repository.EliminarProducto(producto.GetModel())
}
