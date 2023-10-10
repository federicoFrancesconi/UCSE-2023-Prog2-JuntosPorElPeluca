package dto

type Producto struct {
	CodigoProducto int     `json:"codigo_producto"`
	Nombre         string  `json:"nombre"`
	Descripcion    string  `json:"descripcion"`
	Precio         float64 `json:"precio"`
	Stock          int     `json:"stock"`
}
