package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type Camion struct {
	Patente                  string
	PesoMaximo               int
	CostoPorKilometro        float32
	FechaCreacion            time.Time
	FechaUltimaActualizacion time.Time
	IdCreador                int
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
