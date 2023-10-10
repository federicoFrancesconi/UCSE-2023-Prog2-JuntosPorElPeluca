package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	CrearCamion(camion model.Camion) (*mongo.InsertOneResult, error)
	ObtenerCamionPorPatente(patente string) (model.Camion, error)
	ObtenerCamiones() ([]model.Camion, error)
	ActualizarCamion(camion model.Camion) (*mongo.UpdateResult, error)
	EliminarCamion(id primitive.ObjectID) (*mongo.DeleteResult, error)
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

func (repository CamionRepository) CrearCamion(camion model.Camion) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	resultado, err := collection.InsertOne(context.Background(), camion)
	return resultado, err
}

func (repository CamionRepository) ActualizarCamion(camion model.Camion) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": camion.Patente}
	actualizacion := bson.M{"$set": camion}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	return resultado, err
}

func (repository CamionRepository) EliminarCamion(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"_id": id}
	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}
