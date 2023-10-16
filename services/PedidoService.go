package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"time"
)

type PedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
	envioRepository repositories.EnvioRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(pedido *dto.Pedido) error
	ObtenerPedidoPorId(id int) (*dto.Pedido, error)
	ObtenerPedidosFiltrados() ([]dto.Pedido, error)
	ObtenerPesoPedido(id int) (float32, error)
	EnviarPedido(id int) error
	AceptarPedido(id int) error
	CancelarPedido(id int) error
	EntregarPedido(id int) error
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface, envioRepository repositories.EnvioRepository) *PedidoService {
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

func (service *PedidoService) ObtenerPedidos(idEnvio int, estado model.EstadoPedido, fechaCreacionComienzo time.Time, fechaCreacionFin time.Time) ([]dto.Pedido, error) {
	var idPedidos []int
	
	//Lo primero es ver si hace falta filtrar por envio
	if idEnvio != 0 {
		//Buscamos el envio
		envio, err := service.envioRepository.ObtenerEnvioPorId(idEnvio)

		if err != nil {
			return nil, err
		}

		//Si el envio existe, obtenemos la lista de pedidos del mismo
		idPedidos = envio.Pedidos
	}
	
	pedidos, err := service.pedidoRepository.ObtenerPedidosFiltrados(idPedidos, estado, fechaCreacionComienzo, fechaCreacionFin)
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

func (service *PedidoService) ObtenerPedidoPorId(id int) (*dto.Pedido, error) {
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(id)

	if err != nil {
		return nil, err
	}

	pedidoDTO := dto.NewPedido(pedido)

	return pedidoDTO, nil
}

func (service *PedidoService) ObtenerPesoPedido(idPedido int) (float32, error) {
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

	if err != nil {
		return 0, err
	}

	//Calculo el peso del pedido sumando el peso de cada producto elegido
	var peso float32 = 0
	for _, producto := range pedido.ProductosElegidos {
		peso += producto.ObtenerPesoProductoPedido()
	}

	return peso, nil
}

func (service *PedidoService) EnviarPedido(idPedido int) error {
	//Primero buscamos el pedido a enviar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Aceptado
	if pedido.Estado != model.Aceptado {
		return nil
	}

	//Cambia el estado del pedido a Para enviar, si es que no estaba ya en ese estado
	if pedido.Estado != model.ParaEnviar {
		pedido.Estado = model.ParaEnviar
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}

func (service *PedidoService) AceptarPedido(idPedido int) error {
	//Primero buscamos el pedido a aceptar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return nil
	}

	//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Aceptado {
		pedido.Estado = model.Aceptado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}

func (service *PedidoService) CancelarPedido(idPedido int) error {
	//Primero buscamos el pedido a cancelar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

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

func (service *PedidoService) EntregarPedido(idPedido int) error {
	//Primero buscamos el pedido a entregar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Para enviar
	if pedido.Estado != model.ParaEnviar {
		return nil
	}

	//Cambia el estado del pedido a Enviado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Enviado {
		pedido.Estado = model.Enviado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)

	//Descuenta el stock de los productos
}
