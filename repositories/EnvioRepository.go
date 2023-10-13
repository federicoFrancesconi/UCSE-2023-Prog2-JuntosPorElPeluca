package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type EnvioRepositoryInterface interface {
	CrearEnvio(envio model.Envio) error
	ObtenerEnvioPorId(id int) (model.Envio, error)
	ObtenerEnvios() ([]model.Envio, error)
	ActualizarEnvio(envio model.Envio) error
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

	if err != nil {
		return nil, err
	}

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

func (repository EnvioRepository) ObtenerEnvioPorId(id int) (model.Envio, error) {
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

func (repository EnvioRepository) CrearEnvio(envio model.Envio) error {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	//Coloco las fechas
	envio.FechaCreacion = time.Now()
	envio.FechaUltimaActualizacion = time.Now()

	_, err := collection.InsertOne(context.Background(), envio)

	return err
}

func (repository EnvioRepository) ActualizarEnvio(envio model.Envio) error {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"id": envio.Id}

	//seteo la fecha de actualizacion
	envio.FechaUltimaActualizacion = time.Now()

	actualizacion := bson.M{"$set": envio}

	_, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	return err
}
