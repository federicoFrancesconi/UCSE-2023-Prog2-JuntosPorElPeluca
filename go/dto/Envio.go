package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"time"
)

type Envio struct {
	Id                       string            `json:"id"`
	FechaCreacion            time.Time         `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time         `json:"fecha_ultima_actualizacion"`
	PatenteCamion            string            `json:"patente_camion"`
	Paradas                  []Parada          `json:"paradas"`
	Pedidos                  []string          `json:"pedidos"`
	IdCreador                string            `json:"id_creador"`
	Estado                   model.EstadoEnvio `json:"estado"`
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		Id:                       utils.GetStringIDFromObjectID(envio.ObjectId),
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
		ObjectId:                 utils.GetObjectIDFromStringID(envio.Id),
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
