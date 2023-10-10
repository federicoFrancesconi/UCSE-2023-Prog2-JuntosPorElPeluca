package dto

import "UCSE-2023-Prog2-TPIntegrador/model/productos"

type Producto struct {
	CodigoProducto int                    `json:"codigo_producto"`
	TipoDeProducto productos.TipoProducto `json:"tipo_producto"`
	Nombre         string                 `json:"nombre"`
	PesoUnitario   float32                `json:"peso_unitario"`
	PrecioUnitario float32                `json:"precio_unitario"`
	StockMinimo    int                    `json:"stock_minimo"`
	StockActual    int                    `json:"stock_actual"`
}

// Crea el dto a partir del modelo
func NewProducto(producto *productos.Producto) *Producto {
	return &Producto{
		CodigoProducto: producto.CodigoProducto,
		TipoDeProducto: producto.TipoDeProducto,
		Nombre:         producto.Nombre,
		PrecioUnitario: producto.PrecioUnitario,
		PesoUnitario:   producto.PesoUnitario,
		StockMinimo:    producto.StockMinimo,
		StockActual:    producto.StockActual,
	}
}

// Crea el modelo a partir del dto
func (producto Producto) GetModel() productos.Producto {
	return productos.Producto{
		CodigoProducto: producto.CodigoProducto,
		TipoDeProducto: producto.TipoDeProducto,
		Nombre:         producto.Nombre,
		PrecioUnitario: producto.PrecioUnitario,
		PesoUnitario:   producto.PesoUnitario,
		StockMinimo:    producto.StockMinimo,
		StockActual:    producto.StockActual,
	}
}
