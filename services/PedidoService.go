package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type PedidoService struct {
	repository repositories.PedidoRepositoryInterface
}

type PedidoServiceInterface interface {
	CrearPedido(pedido *dto.Pedido) error
	ObtenerPedidos() ([]dto.Pedido, error)
	ObtenerPesoPedido(pedido *dto.Pedido) (float32, error)
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
