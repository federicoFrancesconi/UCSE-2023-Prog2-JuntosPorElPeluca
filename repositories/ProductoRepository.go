package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductoRepositoryInterface interface {
	CrearProducto(producto model.Producto) error
	ObtenerProductoPorCodigo(codigoProducto int) (*model.Producto, error)
	obtenerProductos(filtro bson.M) ([]*model.Producto, error)
	ObtenerProductosConStockMenorAlMinimo() ([]*model.Producto, error)
	ObtenerProductosConStockMenorAlMinimoPorTipo(model.TipoProducto) ([]*model.Producto, error)
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
	//Seteamos las fechas del producto
	producto.FechaCreacion = time.Now()
	producto.FechaUltimaActualizacion = time.Now()

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

func (repository *ProductoRepository) obtenerProductos(filtro bson.M) ([]*model.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	var productosList []*model.Producto

	cursor, err := collection.Find(context.Background(), filtro)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var producto model.Producto

		err := cursor.Decode(&producto)

		if err != nil {
			return nil, err
		}

		productosList = append(productosList, &producto)
	}

	if err := cursor.Err(); err != nil {
        return nil, err
    }

	return productosList, nil
}

func (repository *ProductoRepository) ObtenerProductosConStockMenorAlMinimo() ([]*model.Producto, error) {
	//Creamos un filtro que tenga en cuenta que el stock actual sea menor al minimo
	filtro := bson.M{
		"stock_actual": bson.M{"$lt": "$stock_minimo"},
	}

	return repository.obtenerProductos(filtro)
}

func (repository *ProductoRepository) ObtenerProductosConStockMenorAlMinimoPorTipo(tipoProducto model.TipoProducto) ([]*model.Producto, error) {
	//Creamos un filtro que tenga en cuenta tipo de producto, y que el stock actual sea menor al minimo
    filtro := bson.M{
        "tipo_producto": tipoProducto,
        "stock_actual": bson.M{"$lt": "$stock_minimo"},
    }

	return repository.obtenerProductos(filtro)
}

func (repository *ProductoRepository) ActualizarProducto(producto model.Producto) error {
	//Actualizamos la fecha de actualizacion del producto
	producto.FechaUltimaActualizacion = time.Now()

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
