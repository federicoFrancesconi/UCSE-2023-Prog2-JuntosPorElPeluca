package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/model"
	"TPIntegrador/repositories"
	"TPIntegrador/utils"
	"errors"
)

type ProductoService struct {
	repository repositories.ProductoRepositoryInterface
}

type ProductoServiceInterface interface {
	CrearProducto(*dto.Producto, *dto.User) error
	ObtenerProductos(utils.FiltroProducto) ([]dto.Producto, error)
	ActualizarProducto(*dto.Producto, *dto.User) error
	EliminarProducto(*dto.Producto, *dto.User) error
}

func NewProductoService(repository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		repository: repository,
	}
}

func (service *ProductoService) CrearProducto(producto *dto.Producto, usuario *dto.User) error {
	//Valido el tipo de producto
	if !model.EsUnTipoProductoValido(producto.TipoDeProducto) {
		return errors.New("el tipo de producto ingresado no es válido")
	}

	//valido el usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para crear un producto")
	}

	//Le agregamos el codigo del usuario que lo creo
	producto.IdCreador = usuario.Codigo

	return service.repository.CrearProducto(producto.GetModel())
}

func (service *ProductoService) ObtenerProductos(filtro utils.FiltroProducto) ([]dto.Producto, error) {
	//Valido el tipo de producto que usa para filtrar
	if !model.EsUnTipoProductoValido(filtro.TipoProducto) && filtro.TipoProducto != "" {
		return nil, errors.New("el tipo de producto ingresado no es válido")
	}

	productos, err := service.repository.ObtenerProductos(filtro)

	if err != nil {
		return nil, err
	}

	//Inicializamos el slice de productosDTO por si no hay productos
	productosDTO := make([]dto.Producto, 0)

	for _, producto := range productos {
		productosDTO = append(productosDTO, *dto.NewProducto(producto))
	}

	return productosDTO, nil
}

func (service *ProductoService) ActualizarProducto(producto *dto.Producto, usuario *dto.User) error {
	//Valido el tipo de producto
	if !model.EsUnTipoProductoValido(producto.TipoDeProducto) {
		return errors.New("el tipo de producto ingresado no es válido")
	}

	//valido el usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para actualizar un producto")
	}

	return service.repository.ActualizarProducto(producto.GetModel())
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto, usuario *dto.User) error {
	//valido el usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para eliminar un producto")
	}

	return service.repository.EliminarProducto(producto.GetModel())
}

func (service *ProductoService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == string(utils.Administrador)
}
