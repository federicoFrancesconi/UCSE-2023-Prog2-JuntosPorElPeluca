package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type Envio struct {
	Id                       int
	FechaCreacion            time.Time
	FechaUltimaActualizacion time.Time
	PatenteCamion            string
	Paradas                  []Parada
	Pedidos                  []int
	IdCreador                int
	Estado                   model.EstadoEnvio
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		Id:                       envio.Id,
		FechaCreacion:            envio.FechaCreacion,
		FechaUltimaActualizacion: envio.FechaUltimaActualizacion,
		PatenteCamion:            envio.PatenteCamion,
		Paradas:                  NewParadas(envio.Paradas),
		Pedidos:                  envio.Pedidos,
		IdCreador:                envio.IdCreador,
		Estado:                   envio.Estado,
	}
}

func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		Id:                       envio.Id,
		FechaCreacion:            envio.FechaCreacion,
		FechaUltimaActualizacion: envio.FechaUltimaActualizacion,
		PatenteCamion:            envio.PatenteCamion,
		Paradas:                  envio.getParadas(),
		Pedidos:                  envio.Pedidos,
		IdCreador:                envio.IdCreador,
		Estado:                   envio.Estado,
	}
}

// Metodo para convertir una lista de Paradas del dto a una lista de Paradas del modelo
func (envio Envio) getParadas() []model.Parada {
	var paradasEnvio []model.Parada
	for _, parada := range envio.Paradas {
		paradasEnvio = append(paradasEnvio, parada.GetModel())
	}
	return paradasEnvio
}

// Metodo para convertir una lista de Paradas del modelo a una lista de Paradas del dto
func NewParadas(paradas []model.Parada) []Parada {
	var paradasEnvio []Parada
	for _, parada := range paradas {
		paradasEnvio = append(paradasEnvio, *NewParada(&parada))
	}
	return paradasEnvio
}
