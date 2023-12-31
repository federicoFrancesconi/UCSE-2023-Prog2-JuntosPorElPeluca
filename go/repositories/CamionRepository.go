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

type CamionRepositoryInterface interface {
	CrearCamion(*model.Camion) error
	ObtenerCamiones(utils.FiltroCamion) ([]*model.Camion, error)
	ActualizarCamion(*model.Camion) error
}

type CamionRepository struct {
	db database.DB
}

func NewCamionRepository(db database.DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

func (repository CamionRepository) CrearCamion(camion *model.Camion) error {
	//Nos aseguramos de que el Id sea creado por mongo
	camion.ObjectId = primitive.NewObjectID()

	//Seteamos las fechas para el objeto camion
	camion.FechaCreacion = time.Now()
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")
	_, err := collection.InsertOne(context.Background(), camion)
	return err
}

func (repository CamionRepository) obtenerCamiones(filtro bson.M) ([]*model.Camion, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	cursor, err := collection.Find(context.TODO(), filtro)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	//Inicializamos el slice de camiones por si no hay camiones
	camiones := make([]*model.Camion, 0)

	for cursor.Next(context.Background()) {
		var camion model.Camion
		err := cursor.Decode(&camion)
		if err != nil {
			return nil, err
		}

		camiones = append(camiones, &camion)
	}

	return camiones, err
}

func (repository CamionRepository) ObtenerCamiones(filtro utils.FiltroCamion) ([]*model.Camion, error) {
	//Inicializamos el filtro vacio
	filtroBD := bson.M{}

	//Si el filtro tiene una patente, la agregamos al filtro de la BD
	if filtro.Patente != "" {
		filtroBD["patente"] = filtro.Patente
	}

	//Si el filtro tiene la propiedad de EstaActivo, la agregamos al filtro de la BD
	if filtro.FiltrarPorEstaActivo {
		filtroBD["esta_activo"] = filtro.EstaActivo
	}

	return repository.obtenerCamiones(filtroBD)
}

func (repository CamionRepository) ActualizarCamion(camion *model.Camion) error {
	//Actualizamos la fecha de actualizacion del camion
	camion.FechaUltimaActualizacion = time.Now()

	collection := repository.db.GetClient().Database("empresa").Collection("camiones")

	filtro := bson.M{"patente": camion.Patente}

	//Solo actualizamos ciertos campos
	actualizacion := bson.M{"$set": bson.M{
		"patente":                    camion.Patente,
		"peso_maximo":                camion.PesoMaximo,
		"costo_por_kilometro":        camion.CostoPorKilometro,
		"fecha_ultima_actualizacion": camion.FechaUltimaActualizacion,
		"esta_activo":                camion.EstaActivo,
	}}

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
