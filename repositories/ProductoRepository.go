package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model/productos"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductoRepositoryInterface interface {
	CrearProducto(producto *productos.Producto) error
	ObtenerProductoPorCodigo(codigoProducto int) (*productos.Producto, error)
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

func (repository *ProductoRepository) CrearProducto(producto *productos.Producto) error {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")
	_, err := collection.InsertOne(context.Background(), producto)
	return err
}

func (repository *ProductoRepository) ObtenerProductoPorCodigo(codigoProducto int) (*productos.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")
	
	filtro := bson.M{"codigo_producto": codigoProducto}

	var producto productos.Producto
	
	err := collection.FindOne(context.Background(), filtro).Decode(&producto)

	if err != nil {
		return nil, err
	}

	return &producto, err
}

func (repository *ProductoRepository) ObtenerProductos() ([]*productos.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	var productosList []*productos.Producto

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var producto productos.Producto

		err := cursor.Decode(&producto)
		
		if err != nil {
			return nil, err
		}

		productosList = append(productosList, &producto)
	}

	return productosList, err
}