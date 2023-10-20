package model

type ProductoPedido struct {
	CodigoProducto int     `bson:"codigo_producto"`
	Nombre         string  `bson:"nombre_producto"`
	Cantidad       int     `bson:"cantidad"`
	PrecioUnitario float32 `bson:"precio_unitario"`
	PesoUnitario   float32 `bson:"peso_unitario"`
}

func (productoPedido ProductoPedido) ObtenerPesoProductoPedido() float32 {
	return productoPedido.PesoUnitario * float32(productoPedido.Cantidad)
}