package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"time"
)

type Envio struct {
	Id                       string
	FechaCreacion            time.Time
	FechaUltimaActualizacion time.Time
	PatenteCamion            string
	Paradas                  []Parada
	Pedidos                  []Pedido
	IdCreador                int
	Estado                   model.EstadoEnvio
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		Id:                       utils.GetStringIDFromObjectID(envio.ObjectId),
		FechaCreacion:            envio.FechaCreacion,
		FechaUltimaActualizacion: envio.FechaUltimaActualizacion,
		PatenteCamion:            envio.PatenteCamion,
		Paradas:                  NewParadas(envio.Paradas),
		Pedidos:                  NewPedidos(envio.Pedidos),
		IdCreador:                envio.IdCreador,
		Estado:                   envio.Estado,
	}
}

func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		FechaCreacion:            envio.FechaCreacion,
		FechaUltimaActualizacion: envio.FechaUltimaActualizacion,
		PatenteCamion:            envio.PatenteCamion,
		Paradas:                  envio.getParadas(),
		Pedidos:                  envio.getPedidos(),
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

// Metodo para convertir una lista de pedidos del dto a una lista de pedidos del modelo
func (envio Envio) getPedidos() []model.Pedido {
	var pedidosEnvio []model.Pedido
	for _, pedido := range envio.Pedidos {
		pedidosEnvio = append(pedidosEnvio, pedido.GetModel())
	}
	return pedidosEnvio
}

// Metodo para convertir una lista de Pedidos del modelo a una lista de Pedidos del dto
func NewPedidos(pedidos []model.Pedido) []Pedido {
	var pedidosEnvio []Pedido
	for _, pedido := range pedidos {
		pedidosEnvio = append(pedidosEnvio, *NewPedido(&pedido))
	}
	return pedidosEnvio
}
