package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type CamionRepositoryInterface interface {
	CrearCamion(camion model.Camion) error
	ObtenerCamionPorPatente(patente string) (model.Camion, error)
	ObtenerCamiones() ([]model.Camion, error)
	ActualizarCamion(camion model.Camion) error
	EliminarCamion(patente string) error
}

type CamionRepository struct {
	db database.DB
}

func NewCamionRepository(db database.DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

func (repository CamionRepository) ObtenerCamiones() ([]model.Camion, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var camiones []model.Camion

	for cursor.Next(context.Background()) {
		var camion model.Camion
		err := cursor.Decode(&camion)
		if err != nil {
			return nil, err
		}

		camiones = append(camiones, camion)
	}

	return camiones, err
}

func (repository CamionRepository) ObtenerCamionPorPatente(patente string) (model.Camion, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": patente}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var camion model.Camion

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&camion)
		if err != nil {
			return camion, err
		}
	}

	return camion, err
}

func (repository CamionRepository) CrearCamion(camion model.Camion) error {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	_, err := collection.InsertOne(context.Background(), camion)
	return err
}

func (repository CamionRepository) ActualizarCamion(camion model.Camion) error {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": camion.Patente}
	actualizacion := bson.M{"$set": camion}
	_, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	return err
}

func (repository CamionRepository) EliminarCamion(patente string) error {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": patente}
	_, err := collection.DeleteOne(context.Background(), filtro)
	return err
}
