package model

type EstadoEnvio int

const (
	ADespachar EstadoEnvio = iota
	EnRuta
	Despachado
)
