package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PedidoRepositoryInterface interface {
	CrearPedido(model.Pedido, string) error
	ObtenerPedidosFiltrados(utils.FiltroPedido) ([]*model.Pedido, error)
	ObtenerPedidoPorId(model.Pedido) (*model.Pedido, error)
	ObtenerCantidadPedidosPorEstado(estado model.EstadoPedido) (int, error)
	ActualizarPedido(pedido model.Pedido) error
}

type PedidoRepository struct {
	db database.DB
}

func NewPedidoRepository(db database.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}

func (repository *PedidoRepository) CrearPedido(pedido model.Pedido, idCreador string) error {
	//Nos aseguramos de que el Id sea creado por mongo
	pedido.ObjectId = primitive.NewObjectID()

	//Seteamos las fechas para el objeto pedido
	pedido.FechaCreacion = time.Now()
	pedido.FechaUltimaActualizacion = time.Now()
	pedido.IdCreador = idCreador

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

func (repository *PedidoRepository) ObtenerPedidoPorId(pedidoConId model.Pedido) (*model.Pedido, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"_id": pedidoConId.ObjectId}

	var pedido model.Pedido

	err := collection.FindOne(context.Background(), filtro).Decode(&pedido)

	if err != nil {
		return nil, err
	}

	return &pedido, err
}

// Falta lo de idEnvio
func (repository *PedidoRepository) ObtenerPedidosFiltrados(filtroEnvio utils.FiltroPedido) ([]*model.Pedido, error) {
	//Desestructuramos el filtro
	idPedidos := filtroEnvio.IdPedidos
	estado := filtroEnvio.Estado
	fechaCreacionComienzo := filtroEnvio.FechaCreacionComienzo
	fechaCreacionFin := filtroEnvio.FechaCreacionFin

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

func (repository *PedidoRepository) ObtenerCantidadPedidosPorEstado(estado model.EstadoPedido) (int, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)

	if err != nil {
		return 0, err
	}

	return int(cantidad), nil
}

func (repository *PedidoRepository) ActualizarPedido(pedido model.Pedido) error {
	pedido.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"_id": pedido.ObjectId}

	actualizacion := bson.M{"$set": pedido}

	operacion, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el pedido a actualizar")
	}

	return err
}
