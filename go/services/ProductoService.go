package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/model"
	"TPIntegrador/repositories"
	"TPIntegrador/utils"
	"errors"
)

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
}

type ProductoServiceInterface interface {
	CrearProducto(*dto.Producto, *dto.User) error
	ObtenerProductos(utils.FiltroProducto) ([]dto.Producto, error)
	ObtenerProductoPorCodigo(*dto.Producto, *dto.User) (*dto.Producto, error)
	ActualizarProducto(*dto.Producto, *dto.User) error
	EliminarProducto(*dto.Producto, *dto.User) error
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface) *ProductoService {
	return &ProductoService{
		productoRepository: productoRepository,
		pedidoRepository:   pedidoRepository,
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

	return service.productoRepository.CrearProducto(producto.GetModel())
}

func (service *ProductoService) ObtenerProductos(filtro utils.FiltroProducto) ([]dto.Producto, error) {
	//Valido el tipo de producto que usa para filtrar
	if !model.EsUnTipoProductoValido(filtro.TipoProducto) && filtro.TipoProducto != "" {
		return nil, errors.New("el tipo de producto ingresado no es válido")
	}

	productos, err := service.productoRepository.ObtenerProductos(filtro)

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

func (service *ProductoService) ObtenerProductoPorCodigo(productoConCodigo *dto.Producto, usuario *dto.User) (*dto.Producto, error) {
	productoDB, err := service.productoRepository.ObtenerProductoPorCodigo(productoConCodigo.GetModel())

	//Inicializamos el envio por si no hay ninguno
	var producto *dto.Producto = &dto.Producto{}

	if err != nil {
		return nil, err
	} else {
		producto = dto.NewProducto(productoDB)
	}

	//valido que el envio sea del camionero que lo esta filtrando
	valido := service.validarRol(usuario)

	if !valido && err != nil {
		return nil, err
	}

	return producto, nil
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

	return service.productoRepository.ActualizarProducto(producto.GetModel())
}

func (service *ProductoService) EliminarProducto(producto *dto.Producto, usuario *dto.User) error {
	//valido el usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para eliminar un producto")
	}

	//Valido que el producto a eliminar no tenga pedidos asociados en estado pendiente o aceptado
	err := service.productoTienePedidosEnCurso(producto)

	if err != nil {
		return err
	}

	return service.productoRepository.EliminarProducto(producto.GetModel())
}

func (service *ProductoService) productoTienePedidosEnCurso(producto *dto.Producto) error {
	//Primero armamos el filtro
	filtroPendientes := utils.FiltroPedido{
		CodigoProducto: producto.CodigoProducto,
		Estado:         model.Pendiente,
	}

	pedidosPendientes, err := service.pedidoRepository.ObtenerPedidos(&filtroPendientes)
	if err != nil {
		return err
	}

	if len(pedidosPendientes) > 0 {
		return errors.New("no se puede eliminar el producto: tiene pedidos pendientes")
	}

	filtroAceptados := utils.FiltroPedido{
		CodigoProducto: producto.CodigoProducto,
		Estado:         model.Aceptado,
	}

	pedidosAceptados, err := service.pedidoRepository.ObtenerPedidos(&filtroAceptados)

	if err != nil {
		return err
	}

	if len(pedidosAceptados) > 0 {
		return errors.New("no se puede eliminar el producto: tiene pedidos aceptados")
	}

	return nil
}

func (service *ProductoService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == string(utils.Administrador)
}
