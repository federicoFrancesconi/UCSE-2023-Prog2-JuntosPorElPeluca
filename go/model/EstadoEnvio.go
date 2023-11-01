package model

type EstadoEnvio string

const (
	ADespachar EstadoEnvio = "A Despachar"
	EnRuta    EstadoEnvio = "En Ruta"
	Despachado EstadoEnvio = "Despachado"
)

func EsUnEstadoEnvioValido(estado string) bool {
	return estado == string(ADespachar) || estado == string(EnRuta) || estado == string(Despachado)
}
