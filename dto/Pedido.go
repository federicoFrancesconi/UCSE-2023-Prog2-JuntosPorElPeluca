package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model/pedidos"
	"UCSE-2023-Prog2-TPIntegrador/model/productos"
	"time"
)

type Pedido struct {
	ProductosElegidos        []ProductoPedido     `json:"productos_elegidos"`
	CiudadDestino            string               `json:"ciudad_destino"`
	Estado                   pedidos.EstadoPedido `json:"estado"`
	FechaCreacion            time.Time            `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time            `json:"fecha_ultima_actualizacion"`
	IdCreador                int                  `json:"id_creador"`
}

// Metodo para obtener el modelo a partir del dto
func (pedido Pedido) GetModel() pedidos.Pedido {
	return pedidos.Pedido{
		ProductosElegidos:        pedido.getProductosElegidos(),
		CiudadDestino:            pedido.CiudadDestino,
		Estado:                   pedido.Estado,
		FechaCreacion:            pedido.FechaCreacion,
		FechaUltimaActualizacion: pedido.FechaUltimaActualizacion,
		IdCreador:                pedido.IdCreador,
	}
}

// Metodo para crear un dto a partir del modelo
func NewPedido(pedido *pedidos.Pedido) *Pedido {
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
func (pedido Pedido) getProductosElegidos() []productos.ProductoPedido {
	var productosElegidos []productos.ProductoPedido
	for _, producto := range pedido.ProductosElegidos {
		productosElegidos = append(productosElegidos, producto.GetModel())
	}
	return productosElegidos
}

// Metodo para convertir una lista de ProductoPedido del modelo a una lista de ProductoPedido del dto
func NewProductosPedido(productosElegidos []productos.ProductoPedido) []ProductoPedido {
	var productosElegidosDto []ProductoPedido
	for _, producto := range productosElegidos {
		productosElegidosDto = append(productosElegidosDto, *NewProductoPedido(&producto))
	}
	return productosElegidosDto
}
