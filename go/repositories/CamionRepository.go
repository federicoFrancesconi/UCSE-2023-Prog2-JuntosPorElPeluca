package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CamionRepositoryInterface interface {
	CrearCamion(model.Camion) error
	ObtenerCamionPorPatente(model.Camion) (model.Camion, error)
	ObtenerCamiones() ([]model.Camion, error)
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

func (repository CamionRepository) CrearCamion(camion model.Camion) error {
	//Nos aseguramos de que el Id sea creado por mongo
	camion.ObjectId = primitive.NewObjectID()

	//Seteamos las fechas para el objeto camion
	camion.FechaCreacion = time.Now()
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	_, err := collection.InsertOne(context.Background(), camion)
	return err
}

func (repository CamionRepository) obtenerCamiones(filtro bson.M) ([]model.Camion, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	cursor, err := collection.Find(context.TODO(), filtro)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	//Inicializamos el slice de camiones por si no hay camiones
	camiones := make([]model.Camion, 0)

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

func (repository CamionRepository) ObtenerCamiones() ([]model.Camion, error) {
	//Uso un filtro vacio para que no filtre y traiga todos los camiones
	filtroVacio := bson.M{}

	return repository.obtenerCamiones(filtroVacio)
}

func (repository CamionRepository) ObtenerCamionPorPatente(camion model.Camion) (model.Camion, error) {
	filtro := bson.M{"patente": camion.Patente}

	camiones, err := repository.obtenerCamiones(filtro)

	if err != nil {
		return model.Camion{}, err
	}

	//Contempla que no se haya encontrado el camion en la base de datos
	if len(camiones) == 0 {
		return model.Camion{}, nil
	} else {
		return camiones[0], nil
	}
}

func (repository CamionRepository) ActualizarCamion(camion model.Camion) error {
	//Actualizamos la fecha de actualizacion del camion
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	filtro := bson.M{"patente": camion.Patente}

	actualizacion := bson.M{"$set": camion}

	operacion, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)

	if err != nil {
		return err
	}

	//Si no se actualizo ningun camion, devolvemos un error
	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el camion a actualizar")
	}

	return nil
}

func (repository CamionRepository) EliminarCamion(camion model.Camion) error {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	//Generamos el filtro para eliminar el camion
	filtro := bson.M{"patente": camion.Patente}

	_, err := collection.DeleteOne(context.Background(), filtro)

	return err
}
