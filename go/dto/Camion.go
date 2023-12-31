package dto

import (
	"TPIntegrador/model"
	"time"
)

type Camion struct {
	Patente                  string    `json:"patente"`
	PesoMaximo               int       `json:"peso_maximo"`
	CostoPorKilometro        float64   `json:"costo_por_kilometro"`
	FechaCreacion            time.Time `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time `json:"fecha_ultima_actualizacion"`
	IdCreador                string    `json:"id_creador"`
	EstaActivo               bool      `json:"esta_activo"`
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		Patente:                  camion.Patente,
		PesoMaximo:               camion.PesoMaximo,
		CostoPorKilometro:        camion.CostoPorKilometro,
		FechaCreacion:            camion.FechaCreacion,
		FechaUltimaActualizacion: camion.FechaUltimaActualizacion,
		IdCreador:                camion.IdCreador,
		EstaActivo:               camion.EstaActivo,
	}
}

func (camion Camion) GetModel() *model.Camion {
	return &model.Camion{
		Patente:                  camion.Patente,
		PesoMaximo:               camion.PesoMaximo,
		CostoPorKilometro:        camion.CostoPorKilometro,
		FechaCreacion:            camion.FechaCreacion,
		FechaUltimaActualizacion: camion.FechaUltimaActualizacion,
		IdCreador:                camion.IdCreador,
		EstaActivo:               camion.EstaActivo,
	}
}
