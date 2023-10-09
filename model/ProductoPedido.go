package model

type ProductoPedido struct {
	CodigoProducto int     `bson:"codigoProducto" json:"codigoProducto,omitempty"`
	Nombre         string  `bson:"nombreProducto" json:"nombreProducto,omitempty"`
	Cantidad       int     `bson:"cantidad"`
	PrecioUnitario float32 `bson:"precioUnitario"`
	PesoUnitario   float32 `bson:"pesoUnitario"`
}
