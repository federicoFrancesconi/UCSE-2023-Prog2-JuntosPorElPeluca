package dto

import (
	"TPIntegrador/model"
	"TPIntegrador/utils"
	"time"
)

type Pedido struct {
	Id                       string             `json:"id"`
	ProductosElegidos        []ProductoPedido   `json:"productos_elegidos"`
	CiudadDestino            string             `json:"ciudad_destino"`
	Estado                   model.EstadoPedido `json:"estado"`
	FechaCreacion            time.Time          `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `json:"fecha_ultima_actualizacion"`
	IdCreador                string             `json:"id_creador"`
}

// Metodo para obtener el modelo a partir del dto
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		ObjectId:                 utils.GetObjectIDFromStringID(pedido.Id),
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
		Id:                       utils.GetStringIDFromObjectID(pedido.ObjectId),
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

func (pedido Pedido) ObtenerPecioTotal() float32 {
	var precioTotal float32 = 0
	for _, producto := range pedido.ProductosElegidos {
		precioTotal += producto.PrecioUnitario * float32(producto.Cantidad)
	}
	return precioTotal
}
