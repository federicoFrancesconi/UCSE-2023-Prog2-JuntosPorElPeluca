package repositories

import (
	"TPIntegrador/database"
	"TPIntegrador/model"
	"TPIntegrador/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PedidoRepositoryInterface interface {
	CrearPedido(*model.Pedido) error
	ObtenerPedidos(*utils.FiltroPedido) ([]*model.Pedido, error)
	ObtenerPedidoPorId(*model.Pedido) (*model.Pedido, error)
	ObtenerCantidadPedidosPorEstado(model.EstadoPedido) (int, error)
	ActualizarPedido(*model.Pedido) error
}

type PedidoRepository struct {
	db database.DB
}

func NewPedidoRepository(db database.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}

func (repository *PedidoRepository) CrearPedido(pedido *model.Pedido) error {
	//Nos aseguramos de que el Id sea creado por mongo
	pedido.ObjectId = primitive.NewObjectID()

	//Seteamos las fechas para el objeto pedido
	pedido.FechaCreacion = time.Now()
	pedido.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	_, err := collection.InsertOne(context.Background(), pedido)
	return err
}

func (repository *PedidoRepository) obtenerPedidos(filtro bson.M) ([]*model.Pedido, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	cursor, err := collection.Find(context.Background(), filtro)

	if err != nil {
		return nil, err
	}

	//Inicializamos el slice de pedidos por si no hay pedidos
	pedidos := make([]*model.Pedido, 0)

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		err := cursor.Decode(&pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, &pedido)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return pedidos, nil
}

func (repository *PedidoRepository) ObtenerPedidos(filtro *utils.FiltroPedido) ([]*model.Pedido, error) {
	//Desestructuramos el filtro
	idPedidos := filtro.IdPedidos
	codigoProducto := filtro.CodigoProducto
	estado := filtro.Estado
	fechaCreacionComienzo := filtro.FechaCreacionComienzo
	fechaCreacionFin := filtro.FechaCreacionFin

	filter := bson.M{}

	idPedidosObjectIds := make([]primitive.ObjectID, len(idPedidos))

	//Convertimos la lista de idPedidos en una lista de ObjectId
	for i, idPedido := range idPedidos {
		idPedidoObjectId := utils.GetObjectIDFromStringID(idPedido)
		idPedidosObjectIds[i] = idPedidoObjectId
	}

	// Agrego filtros segun los parametros que se pasen
	//Si el array de idPedidos es mayor a 0, uso el filtro, sino no
	if len(idPedidos) > 0 {
		filter["_id"] = bson.M{"$in": idPedidosObjectIds}
	}

	//Tomo el estado vacio como la ausencia de filtro
	if estado != "" {
		filter["estado"] = estado
	}

	//Tomo el codigoProducto vacio como la ausencia de filtro
	if codigoProducto != "" {
		filter["productos_elegidos"] = bson.M{
			"$elemMatch": bson.M{"codigo_producto": codigoProducto},
		}
	}

	//Tomo la fecha de creacion en 0001-01-01 como la ausencia de filtro
	if !fechaCreacionComienzo.IsZero() || !fechaCreacionFin.IsZero() {
		filtroFecha := bson.M{}
		if !fechaCreacionComienzo.IsZero() {
			filtroFecha["$gte"] = fechaCreacionComienzo
		}
		if !fechaCreacionFin.IsZero() {
			filtroFecha["$lte"] = fechaCreacionFin
		}
		filter["fecha_creacion"] = filtroFecha
	}

	return repository.obtenerPedidos(filter)
}

func (repository *PedidoRepository) ObtenerPedidoPorId(pedidoConId *model.Pedido) (*model.Pedido, error) {
	filtro := bson.M{"_id": pedidoConId.ObjectId}

	pedidos, err := repository.obtenerPedidos(filtro)

	if err != nil {
		return nil, err
	}

	//Controlo que la lista este vacia
	if len(pedidos) == 0 {
		return nil, nil
	}

	return pedidos[0], err
}

func (repository *PedidoRepository) ObtenerCantidadPedidosPorEstado(estado model.EstadoPedido) (int, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)

	if err != nil {
		return 0, err
	}

	return int(cantidad), nil
}

func (repository *PedidoRepository) ActualizarPedido(pedido *model.Pedido) error {
	pedido.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"_id": pedido.ObjectId}

	//Solo permitimos que se actualicen ciertos campos
	actualizacion := bson.M{"$set": bson.M{
		"estado":                     pedido.Estado,
		"fecha_ultima_actualizacion": pedido.FechaUltimaActualizacion,
	}}

	operacion, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el pedido a actualizar")
	}

	return err
}
