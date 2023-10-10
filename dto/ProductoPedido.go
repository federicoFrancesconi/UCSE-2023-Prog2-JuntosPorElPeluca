package dto

import (
	"UCSE-2023-Prog2-TPIntegrador/model/productos"
)

type ProductoPedido struct {
	CodigoProducto int     `json:"codigo_producto"`
	Nombre         string  `json:"nombre_producto"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float32 `json:"precio_unitario"`
	PesoUnitario   float32 `json:"peso_unitario"`
}

// Metodo que sirve para crear un ProductoPedido para un pedido
func NewProductoPedidoFromProducto(producto *productos.Producto, cantidad int) *ProductoPedido {
	return &ProductoPedido{
		CodigoProducto: producto.CodigoProducto,
		Nombre:         producto.Nombre,
		Cantidad:       cantidad,
		PrecioUnitario: producto.PrecioUnitario,
		PesoUnitario:   producto.PesoUnitario,
	}
}

// Crea el dto a partir del model
func NewProductoPedido(productoPedido *productos.ProductoPedido) *ProductoPedido {
	return &ProductoPedido{
		CodigoProducto: productoPedido.CodigoProducto,
		Nombre:         productoPedido.Nombre,
		Cantidad:       productoPedido.Cantidad,
		PrecioUnitario: productoPedido.PrecioUnitario,
		PesoUnitario:   productoPedido.PesoUnitario,
	}
}

func (productoPedido ProductoPedido) GetModel() productos.ProductoPedido {
	return productos.ProductoPedido{
		CodigoProducto: productoPedido.CodigoProducto,
		Nombre:         productoPedido.Nombre,
		Cantidad:       productoPedido.Cantidad,
		PrecioUnitario: productoPedido.PrecioUnitario,
		PesoUnitario:   productoPedido.PesoUnitario,
	}
}
