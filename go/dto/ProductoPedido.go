package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/utils"
)

type ProductoPedido struct {
	CodigoProducto string  `json:"codigo_producto"`
	Nombre         string  `json:"nombre_producto"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float32 `json:"precio_unitario"`
	PesoUnitario   float32 `json:"peso_unitario"`
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
