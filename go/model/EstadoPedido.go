package model

type EstadoPedido string

const (
	Pendiente  EstadoPedido = "Pendiente"
	Aceptado   EstadoPedido = "Aceptado"
	Cancelado  EstadoPedido = "Cancelado"
	ParaEnviar EstadoPedido = "Para Enviar"
	Enviado    EstadoPedido = "Enviado"
)

func EsUnEstadoPedidoValido(estado string) bool {
	return estado == string(Pendiente) || estado == string(Aceptado) || estado == string(Cancelado) || estado == string(ParaEnviar) || estado == string(Enviado)
}
