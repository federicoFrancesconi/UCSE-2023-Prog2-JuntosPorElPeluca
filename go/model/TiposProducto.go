package model

type TipoProducto string

const (
	Golosinas    TipoProducto = "Golosinas"
	Bebidas      TipoProducto = "Bebidas"
	Cigarrillos  TipoProducto = "Cigarrillos"
	Comestibles  TipoProducto = "Comestibles"
	HigieneSalud TipoProducto = "Higiene y Salud"
)

func EsUnTipoProductoValido(tipo TipoProducto) bool {
	return tipo == Golosinas || tipo == Bebidas || tipo == Cigarrillos || tipo == Comestibles || tipo == HigieneSalud
}
