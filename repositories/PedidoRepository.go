package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
)

type PedidoRepositoryInterface interface {
	CrearPedido(pedido *model.Pedido) error
	ObtenerPedidoPorId(id int) (*model.Pedido, error)
	ObtenerPedidos() ([]*model.Pedido, error)
	ActualizarPedido(pedido *model.Pedido) error
}

type PedidoRepository struct {
	db database.DB
}

func NewPedidoRepository(db database.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}
