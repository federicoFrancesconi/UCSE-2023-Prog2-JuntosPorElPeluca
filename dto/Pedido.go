package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type Pedido struct {
	ProductosElegidos        []ProductoPedido     `json:"productos_elegidos"`
	CiudadDestino            string               `json:"ciudad_destino"`
	Estado                   model.EstadoPedido `json:"estado"`
	FechaCreacion            time.Time            `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time            `json:"fecha_ultima_actualizacion"`
	IdCreador                int                  `json:"id_creador"`
}

// Metodo para obtener el modelo a partir del dto
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		ProductosElegidos:        pedido.getProductosElegidos(),
		CiudadDestino:            pedido.CiudadDestino,
		Estado:                   pedido.Estado,
		FechaCreacion:            pedido.FechaCreacion,
		FechaUltimaActualizacion: pedido.FechaUltimaActualizacion,
		IdCreador:                pedido.IdCreador,
	}
}

// Metodo para crear un dto a partir del modelo
func NewPedido(pedido *model.Pedido) *Pedido {
	return &Pedido{
		ProductosElegidos:        NewProductosPedido(pedido.ProductosElegidos),
		CiudadDestino:            pedido.CiudadDestino,
		Estado:                   pedido.Estado,
		FechaCreacion:            pedido.FechaCreacion,
		FechaUltimaActualizacion: pedido.FechaUltimaActualizacion,
		IdCreador:                pedido.IdCreador,
	}
}

// Metodo para convertir una lista de ProductoPedido del dto a una lista de ProductoPedido del modelo
func (pedido Pedido) getProductosElegidos() []model.ProductoPedido {
	var productosElegidos []model.ProductoPedido
	for _, producto := range pedido.ProductosElegidos {
		productosElegidos = append(productosElegidos, producto.GetModel())
	}
	return productosElegidos
}

// Metodo para convertir una lista de ProductoPedido del modelo a una lista de ProductoPedido del dto
func NewProductosPedido(productosElegidos []model.ProductoPedido) []ProductoPedido {
	var productosElegidosDto []ProductoPedido
	for _, producto := range productosElegidos {
		productosElegidosDto = append(productosElegidosDto, *NewProductoPedido(&producto))
	}
	return productosElegidosDto
}
