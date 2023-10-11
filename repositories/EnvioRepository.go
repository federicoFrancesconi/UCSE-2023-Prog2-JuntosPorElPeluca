package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	CrearEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	ObtenerEnvioPorId(id string) (model.Envio, error)
	ObtenerEnvios() ([]model.Envio, error)
	ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error)
	EliminarEnvio(id string) (*mongo.DeleteResult, error)
}

type EnvioRepository struct {
	db database.DB
}

func NewEnvioRepository(db database.DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

func (repository EnvioRepository) ObtenerEnvios() ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var envios []model.Envio

	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			return nil, err
		}

		envios = append(envios, envio)
	}

	return envios, err
}

func (repository EnvioRepository) ObtenerEnvioPorId(id string) (model.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": id}

	cursor, err := collection.Find(context.TODO(), filtro)

	var envio model.Envio

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&envio)
		if err != nil {
			return envio, err
		}
	}

	return envio, err
}

func (repository EnvioRepository) CrearEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	result, err := collection.InsertOne(context.Background(), envio)

	return result, err
}

func (repository EnvioRepository) ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": envio.Id}

	actualizacion := bson.M{"$set": envio}

	result, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	return result, err
}

func (repository EnvioRepository) EliminarEnvio(id string) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": id}

	result, err := collection.DeleteOne(context.Background(), filtro)

	return result, err
}
