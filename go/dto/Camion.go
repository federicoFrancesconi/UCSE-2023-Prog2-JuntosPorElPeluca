package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type Camion struct {
	Patente                  string    `json:"patente"`
	PesoMaximo               int       `json:"pesoMaximo"`
	CostoPorKilometro        float32   `json:"costoPorKilometro"`
	FechaCreacion            time.Time `json:"fechaCreacion"`
	FechaUltimaActualizacion time.Time `json:"fechaUltimaActualizacion"`
	IdCreador                int       `json:"idCreador"`
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		Patente:                  camion.Patente,
		PesoMaximo:               camion.PesoMaximo,
		CostoPorKilometro:        camion.CostoPorKilometro,
		FechaCreacion:            camion.FechaCreacion,
		FechaUltimaActualizacion: camion.FechaUltimaActualizacion,
		IdCreador:                camion.IdCreador,
	}
}

func (camion Camion) GetModel() model.Camion {
	return model.Camion{
		Patente:                  camion.Patente,
		PesoMaximo:               camion.PesoMaximo,
		CostoPorKilometro:        camion.CostoPorKilometro,
		FechaCreacion:            camion.FechaCreacion,
		FechaUltimaActualizacion: camion.FechaUltimaActualizacion,
		IdCreador:                camion.IdCreador,
	}
}
