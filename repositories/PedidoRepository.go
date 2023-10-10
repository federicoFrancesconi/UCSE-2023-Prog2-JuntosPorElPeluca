package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model/pedidos"
)

type PedidoRepositoryInterface interface {
	CrearPedido(pedido *pedidos.Pedido) error
	ObtenerPedidoPorId(id int) (*pedidos.Pedido, error)
	ObtenerPedidos() ([]*pedidos.Pedido, error)
	ActualizarPedido(pedido *pedidos.Pedido) error
}

type PedidoRepository struct {
	db database.DB
}

func NewPedidoRepository(db database.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}
