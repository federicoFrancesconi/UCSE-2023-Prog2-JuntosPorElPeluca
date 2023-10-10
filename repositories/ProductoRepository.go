package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model/productos"
)

type ProductoRepositoryInterface interface {
	CrearProducto(producto *productos.Producto) error
	ObtenerProductoPorId(id int) (*productos.Producto, error)
	ObtenerProductos() ([]*productos.Producto, error)
	ActualizarProducto(producto *productos.Producto) error
	EliminarProducto(id int) error
}

type ProductoRepository struct {
	db database.DB
}

func NewProductoRepository(db database.DB) *ProductoRepository {
	return &ProductoRepository{
		db: db,
	}
}