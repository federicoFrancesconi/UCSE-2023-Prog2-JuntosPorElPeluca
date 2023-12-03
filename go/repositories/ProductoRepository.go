package repositories

import (
	"TPIntegrador/database"
	"TPIntegrador/utils"
	"TPIntegrador/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductoRepositoryInterface interface {
	CrearProducto(*model.Producto) error
	ObtenerProductos(utils.FiltroProducto) ([]*model.Producto, error)
	ObtenerProductoPorCodigo(*model.Producto) (*model.Producto, error)
	ActualizarProducto(*model.Producto) error
	EliminarProducto(*model.Producto) error
}

type ProductoRepository struct {
	db database.DB
}

func NewProductoRepository(db database.DB) *ProductoRepository {
	return &ProductoRepository{
		db: db,
	}
}

func (repository *ProductoRepository) CrearProducto(producto *model.Producto) error {
	//Nos aseguramos de que el Id sea creado por mongo
	producto.ObjectId = primitive.NewObjectID()

	//Seteamos las fechas del producto
	producto.FechaCreacion = time.Now()
	producto.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("productos")
	_, err := collection.InsertOne(context.Background(), producto)
	return err
}

func (repository *ProductoRepository) ObtenerProductoPorCodigo(productoConCodigo *model.Producto) (*model.Producto, error) {
	filtro := bson.M{"_id": productoConCodigo.ObjectId}

	productos, err := repository.obtenerProductos(filtro)

	if err != nil {
		return nil, err
	}

	if len(productos) == 0 {
		return nil, errors.New("no se encontró el producto")
	}

	return productos[0], err
}

func (repository *ProductoRepository) ObtenerProductos(filtroProducto utils.FiltroProducto) ([]*model.Producto, error) {
	//Primero creamos el filtro vacio
	filtroDB := bson.M{}

	// Write a custom JavaScript expression using $where
	customJavaScript := "this.stock_actual < this.stock_minimo"

	//Si quiere filtrar por stock minimo, lo agregamos al filtro
	if filtroProducto.FiltrarPorStockMinimo {
		filtroDB["$where"] = customJavaScript
	}

	//Si quiere filtrar por tipo de producto, lo agregamos al filtro
	if filtroProducto.TipoProducto != "" {
		filtroDB["tipo_producto"] = filtroProducto.TipoProducto
	}

	return repository.obtenerProductos(filtroDB)
}

func (repository *ProductoRepository) obtenerProductos(filtro bson.M) ([]*model.Producto, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	//Inicializamos el slice de productos por si no hay productos
	productosList := make([]*model.Producto, 0)

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

func (repository *ProductoRepository) ActualizarProducto(producto *model.Producto) error {
	//Actualizamos la fecha de actualizacion del producto
	producto.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	filtro := bson.M{"_id": producto.ObjectId}

	//Creo una operacion personalizada, para que no actualice nunca la fecha de creacion o el id del creador
	actualizacion := bson.M{
		"$set": bson.M{
			"tipo_producto":   producto.TipoDeProducto,
			"nombre":          producto.Nombre,
			"peso_unitario":   producto.PesoUnitario,
			"precio_unitario": producto.PrecioUnitario,
			"stock_minimo":    producto.StockMinimo,
			"stock_actual":    producto.StockActual,
		},
	}

	operacion, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el producto a actualizar")
	}

	return err
}

func (repository *ProductoRepository) EliminarProducto(producto *model.Producto) error {
	collection := repository.db.GetClient().Database("empresa").Collection("productos")

	filtro := bson.M{"_id": producto.ObjectId}

	_, err := collection.DeleteOne(context.Background(), filtro)

	return err
}
