package pedidos

type EstadoPedido int

const (
	Pendiente EstadoPedido = iota
	Aceptado
	Cancelado
	ParaEnviar
	Enviado
)
