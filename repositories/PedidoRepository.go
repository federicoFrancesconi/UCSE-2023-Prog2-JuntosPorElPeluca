package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type PedidoRepositoryInterface interface {
	CrearPedido(pedido model.Pedido) error
	ObtenerPedidoPorId(id int) (*model.Pedido, error)
	ObtenerPedidos() ([]*model.Pedido, error)
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

func (repository *PedidoRepository) CrearPedido(pedido model.Pedido) error {
	pedido.FechaCreacion = time.Now()
	pedido.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")
	_, err := collection.InsertOne(context.Background(), pedido)
	return err
}

func (repository *PedidoRepository) ObtenerPedidoPorId(id int) (*model.Pedido, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"id": id}

	var pedido model.Pedido

	err := collection.FindOne(context.Background(), filtro).Decode(&pedido)

	if err != nil {
		return nil, err
	}

	return &pedido, err
}

func (repository *PedidoRepository) ObtenerPedidos() ([]*model.Pedido, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	//Creamos un filtro vacío para que traiga todos los pedidos
	filtro := bson.M{}

	cursor, err := collection.Find(context.Background(), filtro)

	if err != nil {
		return nil, err
	}

	var pedidos []*model.Pedido

	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		err := cursor.Decode(&pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, &pedido)
	}

	return pedidos, err
}

func (repository *PedidoRepository) ActualizarPedido(pedido model.Pedido) error {
	pedido.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("pedidos")

	filtro := bson.M{"id": pedido.Id}

	actualizacion := bson.M{"$set": pedido}

	_, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	return err
}
