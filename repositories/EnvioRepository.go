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
	ObtenerEnvios(filtro bson.M) ([]model.Envio, error)
	ObtenerEnviosFiltrados(patente string, estado model.EstadoEnvio, ultimaParada string, fechaCreacionComienzo time.Time, fechaCreacionFin time.Time) ([]model.Envio, error)
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

func (repository EnvioRepository) ObtenerEnvios(filtro bson.M) ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

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

func (repository EnvioRepository) ObtenerEnviosFiltrados(patente string, estado model.EstadoEnvio, ultimaParada string, fechaCreacionComienzo time.Time, fechaCreacionFin time.Time) ([]model.Envio, error) {
	filtro := bson.M{}

	//Solo filtra por patente si le pasamos un valor distinto de ""
	if patente != "" {
		filtro["patente_camion"] = patente
	}

	//Solo filtra por estado si le pasamos un estado positivo
	if estado != (-1) {
		filtro["estado"] = estado
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
		filtro["fecha_creacion"] = filtroFecha
	}

	//TODO: hay que probar esta parte
	if ultimaParada != "" {
		if ultimaParada != "" {
			filtro["paradas"] = bson.M{
				"$elemMatch": bson.M{
					"ciudad": ultimaParada,
				},
			}
			filtro["paradas.$slice"] = -1
		}
	}

	return repository.ObtenerEnvios(filtro)
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
