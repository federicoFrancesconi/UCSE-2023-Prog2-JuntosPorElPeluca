package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
)

type PedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
	envioRepository repositories.EnvioRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(*dto.Pedido) error
	ObtenerPedidoPorId(*dto.Pedido) (*dto.Pedido, error)
	ObtenerPedidosFiltrados(utils.FiltroPedido) ([]dto.Pedido, error)
	AceptarPedido(*dto.Pedido) error
	CancelarPedido(*dto.Pedido) error
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface, envioRepository repositories.EnvioRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository: pedidoRepository,
		envioRepository: envioRepository,
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

	var idPedidos []int
	
	//Lo primero es ver si hace falta filtrar por envio
	if idEnvio != 0 {
		//Buscamos el envio
		envio, err := service.envioRepository.ObtenerEnvioPorId(model.Envio{Id: idEnvio})

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

//TODO: falta verificar que se tenga el stock disponible antes de aceptarlo
func (service *PedidoService) AceptarPedido(pedidoPorAceptar *dto.Pedido) error {
	//Primero buscamos el pedido a aceptar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorAceptar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return nil
	}

	//Verifica que haya stock disponible para aceptar el pedido
	

	//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Aceptado {
		pedido.Estado = model.Aceptado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
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
