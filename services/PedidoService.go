package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type PedidoService struct {
	repository repositories.PedidoRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(pedido *dto.Pedido) error
	ObtenerPedidos() ([]dto.Pedido, error)
	ObtenerPesoPedido(id int) (float32, error)
	CabeEnCamion(pedido *dto.Pedido) (bool, error)
	EnviarPedido(pedido *dto.Pedido) error
	AceptarPedido(pedido *dto.Pedido) error
	CancelarPedido(pedido *dto.Pedido) error
	EntregarPedido(pedido *dto.Pedido) error
}

func NewPedidoService(repository repositories.PedidoRepositoryInterface) *PedidoService {
	return &PedidoService{
		repository: repository,
	}
}

func (service *PedidoService) CrearPedido(pedido *dto.Pedido) error {
	//Obligamos a que el estado del pedido sea Pendiente
	if pedido.Estado != model.Pendiente {
		pedido.Estado = model.Pendiente
	}

	return service.repository.CrearPedido(pedido.GetModel())
}

func (service *PedidoService) ObtenerPedidos() ([]dto.Pedido, error) {
	pedidos, err := service.repository.ObtenerPedidos()
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

func (service *PedidoService) obtenerPesoPedido(id int) (float32, error) {
	pedido, err := service.repository.ObtenerPedidoPorId(id)

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

//Charlar como hacer esto
// func (service *PedidoService) CabeEnCamion(pedido *dto.Pedido, patente string) (bool, error) {
// 	pesoPedido, err := service.obtenerPesoPedido(pedido.Id)

// 	return service.repository.CabeEnCamion(pedido.GetModel())
// }

func (service *PedidoService) EnviarPedido(pedido *dto.Pedido) error {
	//Valida que el pedido esté en estado Aceptado
	if pedido.Estado != model.Aceptado {
		return nil
	}

	//Cambia el estado del pedido a Para enviar, si es que no estaba ya en ese estado
	if pedido.Estado != model.ParaEnviar {
		pedido.Estado = model.ParaEnviar
	}

	//Actualiza el pedido en la base de datos
	return service.repository.ActualizarPedido(pedido.GetModel())
}

func (service *PedidoService) AceptarPedido(pedido *dto.Pedido) error {
	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return nil
	}

	//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Aceptado {
		pedido.Estado = model.Aceptado
	}

	//Actualiza el pedido en la base de datos
	return service.repository.ActualizarPedido(pedido.GetModel())
}

func (service *PedidoService) CancelarPedido(pedido *dto.Pedido) error {
	//Valida que el pedido esté en estado Pendiente
	if pedido.Estado != model.Pendiente {
		return nil
	}

	//Cambia el estado del pedido a Cancelado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Cancelado {
		pedido.Estado = model.Cancelado
	}

	//Actualiza el pedido en la base de datos
	return service.repository.ActualizarPedido(pedido.GetModel())
}

func (service *PedidoService) EntregarPedido(pedido *dto.Pedido) error {
	//Valida que el pedido esté en estado Para enviar
	if pedido.Estado != model.ParaEnviar {
		return nil
	}

	//Cambia el estado del pedido a Enviado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Enviado {
		pedido.Estado = model.Enviado
	}

	//Actualiza el pedido en la base de datos
	return service.repository.ActualizarPedido(pedido.GetModel())

	//Descuenta el stock de los productos
}
