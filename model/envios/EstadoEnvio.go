package envios

type EstadoEnvio int

const (
	ADespachar EstadoEnvio = iota
	EnRuta
	Despachado
)
