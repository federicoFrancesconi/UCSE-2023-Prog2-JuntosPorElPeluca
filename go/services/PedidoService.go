package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/model"
	"TPIntegrador/repositories"
	"TPIntegrador/utils"
	"errors"
)

type PedidoService struct {
	pedidoRepository   repositories.PedidoRepositoryInterface
	envioRepository    repositories.EnvioRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(*dto.Pedido, *dto.User) error
	ObtenerPedidos(utils.FiltroPedido) ([]*dto.Pedido, error)
	ObtenerPedidoPorId(*dto.Pedido) (*dto.Pedido, error)
	ObtenerCantidadPedidosPorEstado() ([]utils.CantidadEstado, error)
	AceptarPedido(*dto.Pedido, *dto.User) error
	CancelarPedido(*dto.Pedido, *dto.User) error
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface, envioRepository repositories.EnvioRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository:   pedidoRepository,
		envioRepository:    envioRepository,
		productoRepository: productoRepository,
	}
}

func (service *PedidoService) CrearPedido(pedido *dto.Pedido, usuario *dto.User) error {
	//Validamos el rol del usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para crear un pedido")
	}

	//Aseguramos que el pedido tenga productos
	if len(pedido.ProductosElegidos) == 0 {
		return errors.New("el pedido debe tener al menos un producto")
	}

	//Aseguramos que el pedido tenga destino
	if pedido.CiudadDestino == "" {
		return errors.New("el pedido debe tener un destino")
	}

	//Obligamos a que el estado del pedido sea Pendiente
	if pedido.Estado != model.Pendiente {
		pedido.Estado = model.Pendiente
	}

	//Le agregamos el codigo del usuario que lo creo
	pedido.IdCreador = usuario.Codigo

	return service.pedidoRepository.CrearPedido(pedido.GetModel())
}

func (service *PedidoService) ObtenerPedidos(filtroPedido utils.FiltroPedido) ([]*dto.Pedido, error) {
	//Obtenemos el id del envio, si es que se filtró por el mismo
	idEnvio := filtroPedido.IdEnvio

	var idPedidos []string

	//Lo primero es ver si hace falta filtrar por envio
	if idEnvio != "" {
		//Generamos el envio para buscar
		envioParaBuscar := dto.Envio{Id: idEnvio}

		//Buscamos el envio
		envio, err := service.envioRepository.ObtenerEnvioPorId(envioParaBuscar.GetModel())

		if err != nil {
			return nil, err
		}

		//Si el envio existe, obtenemos la lista de pedidos del mismo
		idPedidos = envio.Pedidos

		if idPedidos == nil {
			//Si no hay pedidos, devolvemos un array vacio
			return []*dto.Pedido{}, nil
		}
	}

	//Asignamos la lista de pedidos al filtro
	filtroPedido.IdPedidos = idPedidos

	//Validamos el estado del pedido
	if !model.EsUnEstadoPedidoValido(filtroPedido.Estado) && filtroPedido.Estado != "" {
		return nil, errors.New("el estado ingresado para filtrar no es válido")
	}

	pedidos, err := service.pedidoRepository.ObtenerPedidos(&filtroPedido)
	if err != nil {
		return nil, err
	}

	//Inicializamos el array de pedidosDTO por si esta vacio
	var pedidosDTO []*dto.Pedido = []*dto.Pedido{}

	for _, pedido := range pedidos {
		pedidoDTO := dto.NewPedido(pedido)
		pedidosDTO = append(pedidosDTO, pedidoDTO)
	}

	return pedidosDTO, nil
}

func (service *PedidoService) ObtenerPedidoPorId(pedidoConId *dto.Pedido) (*dto.Pedido, error) {
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoConId.GetModel())

	if err != nil {
		return nil, err
	}

	if pedido == nil {
		return &dto.Pedido{}, nil
	} else {
		pedidoDTO := dto.NewPedido(pedido)
		return pedidoDTO, nil
	}
}

func (service *PedidoService) AceptarPedido(pedidoPorAceptar *dto.Pedido, usuario *dto.User) error {
	//Validamos el rol del usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para aceptar el pedido")
	}

	//Primero buscamos el pedido a aceptar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorAceptar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return errors.New("el pedido no se encuentra en estado Pendiente")
	}

	//Verifica que haya stock disponible para aceptar el pedido
	if !service.hayStockDisponiblePedido(pedido) {
		return errors.New("no hay stock disponible para aceptar el pedido")
	}

	//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Aceptado {
		pedido.Estado = model.Aceptado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(pedido)
}

func (service *PedidoService) hayStockDisponiblePedido(pedido *model.Pedido) bool {
	//Busco los productos del pedido
	productosPedido := pedido.ProductosElegidos

	//Recorro los productos del pedido
	for _, productoPedido := range productosPedido {

		//Armo un objeto producto con el ID para buscar en la base de datos
		dtoProductoParaBuscar := dto.Producto{CodigoProducto: productoPedido.CodigoProducto}

		//Busco el producto en la base de datos
		producto, err := service.productoRepository.ObtenerProductoPorCodigo(dtoProductoParaBuscar.GetModel())

		if err != nil {
			return false
		}

		//Verifico que haya stock disponible para el producto
		if productoPedido.Cantidad > producto.StockActual {
			return false
		}
	}

	//Si finalice el bucle, es porque hay stock de todos los productos
	return true
}

func (service *PedidoService) ObtenerCantidadPedidosPorEstado() ([]utils.CantidadEstado, error) {
	//Por cada estado posible de pedidos, obtengo la cantidad de pedidos en ese estado
	cantidadPedidosPendientes, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(model.Pendiente)

	if err != nil {
		return nil, err
	}

	cantidadPedidosAceptados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(model.Aceptado)

	if err != nil {
		return nil, err
	}

	cantidadPedidosCancelados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(model.Cancelado)

	if err != nil {
		return nil, err
	}

	cantidadPedidosParaEnviar, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(model.ParaEnviar)

	if err != nil {
		return nil, err
	}

	cantidadPedidosEnviados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(model.Enviado)

	if err != nil {
		return nil, err
	}

	//Armo el array de CantidadEstado
	cantidadPedidosPorEstados := []utils.CantidadEstado{
		{Estado: string(model.Pendiente), Cantidad: cantidadPedidosPendientes},
		{Estado: string(model.Aceptado), Cantidad: cantidadPedidosAceptados},
		{Estado: string(model.Cancelado), Cantidad: cantidadPedidosCancelados},
		{Estado: string(model.ParaEnviar), Cantidad: cantidadPedidosParaEnviar},
		{Estado: string(model.Enviado), Cantidad: cantidadPedidosEnviados},
	}

	return cantidadPedidosPorEstados, nil
}

func (service *PedidoService) CancelarPedido(pedidoPorCancelar *dto.Pedido, usuario *dto.User) error {
	//Validamos el rol del usuario
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para crear un pedido")
	}

	//Primero buscamos el pedido a cancelar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorCancelar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return errors.New("el pedido no se encuentra en estado Pendiente")
	}

	//Cambia el estado del pedido a Cancelado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Cancelado {
		pedido.Estado = model.Cancelado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(pedido)
}

// valida el rol del usuario
func (service *PedidoService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == string(utils.Operador)
}
