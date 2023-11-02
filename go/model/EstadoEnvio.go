package model

type EstadoEnvio string

const (
	ADespachar EstadoEnvio = "A Despachar"
	EnRuta     EstadoEnvio = "En Ruta"
	Despachado EstadoEnvio = "Despachado"
)

func EsUnEstadoEnvioValido(estado EstadoEnvio) bool {
	return estado == ADespachar || estado == EnRuta || estado == Despachado
}
