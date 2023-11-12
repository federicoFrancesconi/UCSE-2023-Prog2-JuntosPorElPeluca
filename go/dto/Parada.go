package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
)

type Parada struct {
	IdEnvio      string `json:"id_envio"`
	Ciudad       string `json:"ciudad"`
	KmRecorridos int    `json:"km_recorridos"`
}

// Metodo para obtener el modelo a partir del dto
func (parada Parada) GetModel() model.Parada {
	return model.Parada{
		IdEnvio:      parada.IdEnvio,
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}

// Metodo para crear un dto a partir del modelo
func NewParada(parada *model.Parada) *Parada {
	return &Parada{
		IdEnvio:      parada.IdEnvio,
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}
