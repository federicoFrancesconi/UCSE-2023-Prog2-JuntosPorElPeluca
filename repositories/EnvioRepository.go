package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model/envios"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	CrearEnvio(envio *envios.Envio) (*mongo.InsertOneResult, error)
	ObtenerEnvioPorId(id int) (envios.Envio, error)
	ObtenerEnvios() ([]envios.Envio, error)
	ActualizarEnvio(envio *envios.Envio) (*mongo.UpdateResult, error)
	EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type EnvioRepository struct {
	db database.DB
}

func NewEnvioRepository(db database.DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

func (repository EnvioRepository) ObtenerEnvios() ([]envios.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var listaEnvios []envios.Envio

	for cursor.Next(context.Background()) {
		var envio envios.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			return nil, err
		}

		listaEnvios = append(listaEnvios, envio)
	}

	return listaEnvios, err
}

func (repository EnvioRepository) ObtenerEnvioPorId(id string) (envios.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": id}

	cursor, err := collection.Find(context.TODO(), filtro)

	var envio envios.Envio

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&envio)
		if err != nil {
			return envio, err
		}
	}

	return envio, err
}

func (repository EnvioRepository) CrearEnvio(envio envios.Envio) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	result, err := collection.InsertOne(context.Background(), envio)

	return result, err
}

func (repository EnvioRepository) ActualizarEnvio(envio envios.Envio) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": envio.Id}

	actualizacion := bson.M{"$set": envio}

	result, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	return result, err
}

func (repository EnvioRepository) EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": id}

	result, err := collection.DeleteOne(context.Background(), filtro)

	return result, err
}
