package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductoRepositoryInterface interface {
	CrearProducto(producto model.Producto) error
	ObtenerProductoPorCodigo(codigoProducto int) (*model.Producto, error)
	ObtenerProductos() ([]*model.Producto, error)
	ActualizarProducto(producto model.Producto) error
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

func (repository *ProductoRepository) CrearProducto(producto model.Producto) error {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")
	_, err := collection.InsertOne(context.Background(), producto)
	return err
}

func (repository *ProductoRepository) ObtenerProductoPorCodigo(codigoProducto int) (*model.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	filtro := bson.M{"codigo_producto": codigoProducto}

	var producto model.Producto

	err := collection.FindOne(context.Background(), filtro).Decode(&producto)

	if err != nil {
		return nil, err
	}

	return &producto, err
}

func (repository *ProductoRepository) ObtenerProductos() ([]*model.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	var productosList []*model.Producto

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var producto model.Producto

		err := cursor.Decode(&producto)

		if err != nil {
			return nil, err
		}

		productosList = append(productosList, &producto)
	}

	return productosList, err
}

func (repository *ProductoRepository) ActualizarProducto(producto model.Producto) error {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	filtro := bson.M{"codigo_producto": producto.CodigoProducto}

	//TODO: ver si anda este tipo de set
	_, err := collection.UpdateOne(context.Background(), filtro, bson.M{"$set": producto})

	return err
}

func (repository *ProductoRepository) EliminarProducto(id int) error {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	filtro := bson.M{"codigo_producto": id}

	_, err := collection.DeleteOne(context.Background(), filtro)

	return err
}
