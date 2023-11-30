package model

type ProductoPedido struct {
	CodigoProducto string  `bson:"codigo_producto"`
	Nombre         string  `bson:"nombre_producto"`
	Cantidad       int     `bson:"cantidad"`
	PrecioUnitario float64 `bson:"precio_unitario"`
	PesoUnitario   float64 `bson:"peso_unitario"`
}

func (productoPedido ProductoPedido) ObtenerPesoProductoPedido() float64 {
	return productoPedido.PesoUnitario * float64(productoPedido.Cantidad)
}
