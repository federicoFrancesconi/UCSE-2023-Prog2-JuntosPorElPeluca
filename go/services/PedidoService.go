package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"errors"
)

type PedidoService struct {
	pedidoRepository   repositories.PedidoRepositoryInterface
	envioRepository    repositories.EnvioRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(*dto.Pedido) error
	ObtenerPedidoPorId(*dto.Pedido) (*dto.Pedido, error)
	ObtenerPedidosFiltrados(utils.FiltroPedido) ([]dto.Pedido, error)
	ObtenerCantidadPedidosPorEstado() ([]utils.CantidadEstado, error)
	AceptarPedido(*dto.Pedido) error
	CancelarPedido(*dto.Pedido) error
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface, envioRepository repositories.EnvioRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository:   pedidoRepository,
		envioRepository:    envioRepository,
		productoRepository: productoRepository,
	}
}

func (service *PedidoService) CrearPedido(pedido *dto.Pedido) error {
	//Obligamos a que el estado del pedido sea Pendiente
	if pedido.Estado != model.Pendiente {
		pedido.Estado = model.Pendiente
	}

	return service.pedidoRepository.CrearPedido(pedido.GetModel())
}

func (service *PedidoService) ObtenerPedidosFiltrados(filtroPedido utils.FiltroPedido) ([]dto.Pedido, error) {
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
	}

	//Asignamos la lista de pedidos al filtro
	filtroPedido.IdPedidos = idPedidos

	pedidos, err := service.pedidoRepository.ObtenerPedidosFiltrados(filtroPedido)
	if err != nil {
		return nil, err
	}

	var pedidosDTO []dto.Pedido

	for _, pedido := range pedidos {
		pedidoDTO := *dto.NewPedido(pedido)
		pedidosDTO = append(pedidosDTO, pedidoDTO)
	}

	return pedidosDTO, nil
}

func (service *PedidoService) ObtenerPedidoPorId(pedidoConId *dto.Pedido) (*dto.Pedido, error) {
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoConId.GetModel())

	if err != nil {
		return nil, err
	}

	pedidoDTO := dto.NewPedido(pedido)

	return pedidoDTO, nil
}

func (service *PedidoService) AceptarPedido(pedidoPorAceptar *dto.Pedido) error {
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
	return service.pedidoRepository.ActualizarPedido(*pedido)
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
		{ Estado: "Pendiente", Cantidad: cantidadPedidosPendientes },
		{ Estado: "Aceptado", Cantidad: cantidadPedidosAceptados },
		{ Estado: "Cancelado", Cantidad: cantidadPedidosCancelados },
		{ Estado: "ParaEnviar", Cantidad: cantidadPedidosParaEnviar },
		{ Estado: "Enviado", Cantidad: cantidadPedidosEnviados },
	}

	return cantidadPedidosPorEstados, nil
}

func (service *PedidoService) CancelarPedido(pedidoPorCancelar *dto.Pedido) error {
	//Primero buscamos el pedido a cancelar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorCancelar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return nil
	}

	//Cambia el estado del pedido a Cancelado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Cancelado {
		pedido.Estado = model.Cancelado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}
