package model

type EstadoPedido int

const (
	Pendiente EstadoPedido = iota
	Aceptado
	Cancelado
	ParaEnviar
	Enviado
)
