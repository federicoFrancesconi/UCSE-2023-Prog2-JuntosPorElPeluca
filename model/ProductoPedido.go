package model

type ProductoPedido struct {
	Pedido   Pedido `bson:"pedido"`
	Cantidad int    `bson:"cantidad"`
}
