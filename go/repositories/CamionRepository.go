package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type CamionRepositoryInterface interface {
	CrearCamion(model.Camion) error
	ObtenerCamionPorPatente(model.Camion) (model.Camion, error)
	ObtenerTodosLosCamiones() ([]model.Camion, error)
	ActualizarCamion(model.Camion) error
	EliminarCamion(model.Camion) error
}

type CamionRepository struct {
	db database.DB
}

func NewCamionRepository(db database.DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

func (repository CamionRepository) obtenerCamiones(filtro bson.M) ([]model.Camion, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	cursor, err := collection.Find(context.TODO(), filtro)

	if err != nil {
		return nil, err
	}

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

func (repository CamionRepository) ObtenerTodosLosCamiones() ([]model.Camion, error) {
	//Uso un filtro vacio para que no filtre y traiga todos los camiones
	filtroVacio := bson.M{}

	return repository.obtenerCamiones(filtroVacio)
}

func (repository CamionRepository) ObtenerCamionPorPatente(camion model.Camion) (model.Camion, error) {
	filtro := bson.M{"patente": camion.Patente}

	camiones, err := repository.obtenerCamiones(filtro)

	// if err != nil {
	// 	return model.Camion{}, err
	// }

	return camiones[0], err
}

func (repository CamionRepository) CrearCamion(camion model.Camion) error {
	//Seteamos las fechas para el objeto camion
	camion.FechaCreacion = time.Now()
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	_, err := collection.InsertOne(context.Background(), camion)
	return err
}

func (repository CamionRepository) ActualizarCamion(camion model.Camion) error {
	//Actualizamos la fecha de actualizacion del camion
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": camion.Patente}
	actualizacion := bson.M{"$set": camion}
	_, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	return err
}

func (repository CamionRepository) EliminarCamion(camion model.Camion) error {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	filtro := bson.M{"patente": camion.Patente}
	_, err := collection.DeleteOne(context.Background(), filtro)
	return err
}
