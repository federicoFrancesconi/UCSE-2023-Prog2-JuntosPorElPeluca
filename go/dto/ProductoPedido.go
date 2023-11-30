package dto

import (
	"TPIntegrador/model"
	"TPIntegrador/utils"
)

type ProductoPedido struct {
	CodigoProducto string  `json:"codigo_producto"`
	Nombre         string  `json:"nombre_producto"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
	PesoUnitario   float64 `json:"peso_unitario"`
}

// Metodo que sirve para crear un ProductoPedido para un pedido
func NewProductoPedidoFromProducto(producto *model.Producto, cantidad int) *ProductoPedido {
	return &ProductoPedido{
		CodigoProducto: utils.GetStringIDFromObjectID(producto.ObjectId),
		Nombre:         producto.Nombre,
		Cantidad:       cantidad,
		PrecioUnitario: producto.PrecioUnitario,
		PesoUnitario:   producto.PesoUnitario,
	}
}

// Crea el dto a partir del model
func NewProductoPedido(productoPedido *model.ProductoPedido) *ProductoPedido {
	return &ProductoPedido{
		CodigoProducto: productoPedido.CodigoProducto,
		Nombre:         productoPedido.Nombre,
		Cantidad:       productoPedido.Cantidad,
		PrecioUnitario: productoPedido.PrecioUnitario,
		PesoUnitario:   productoPedido.PesoUnitario,
	}
}

func (productoPedido ProductoPedido) GetModel() model.ProductoPedido {
	return model.ProductoPedido{
		CodigoProducto: productoPedido.CodigoProducto,
		Nombre:         productoPedido.Nombre,
		Cantidad:       productoPedido.Cantidad,
		PrecioUnitario: productoPedido.PrecioUnitario,
		PesoUnitario:   productoPedido.PesoUnitario,
	}
}
