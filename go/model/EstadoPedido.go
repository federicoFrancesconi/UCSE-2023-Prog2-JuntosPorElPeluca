package model

type EstadoPedido string

const (
	Pendiente  EstadoPedido = "Pendiente"
	Aceptado   EstadoPedido = "Aceptado"
	Cancelado  EstadoPedido = "Cancelado"
	ParaEnviar EstadoPedido = "Para Enviar"
	Enviado    EstadoPedido = "Enviado"
)

func EsUnEstadoPedidoValido(estado EstadoPedido) bool {
	return estado == Pendiente || estado == Aceptado || estado == Cancelado || estado == ParaEnviar || estado == Enviado
}
