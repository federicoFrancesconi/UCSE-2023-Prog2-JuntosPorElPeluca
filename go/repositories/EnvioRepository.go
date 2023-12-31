package repositories

import (
	"TPIntegrador/database"
	"TPIntegrador/utils"
	"TPIntegrador/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnvioRepositoryInterface interface {
	CrearEnvio(*model.Envio) error
	ObtenerEnvios(*utils.FiltroEnvio) ([]*model.Envio, error)
	ObtenerEnvioPorId(*model.Envio) (*model.Envio, error)
	ObtenerCantidadEnviosPorEstado(model.EstadoEnvio) (int, error)
	ActualizarEnvio(*model.Envio) error
}

type EnvioRepository struct {
	db database.DB
}

func NewEnvioRepository(db database.DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

func (repository EnvioRepository) CrearEnvio(envio *model.Envio) error {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	//Aseguramos que el id sea creado por mongo
	envio.ObjectId = primitive.NewObjectID()

	//Coloco las fechas
	envio.FechaCreacion = time.Now()
	envio.FechaUltimaActualizacion = time.Now()

	_, err := collection.InsertOne(context.Background(), envio)

	return err
}

func (repository EnvioRepository) obtenerEnvios(filtro bson.M) ([]*model.Envio, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	cursor, err := collection.Find(context.TODO(), filtro)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	//Inicializo el slice de envios por si no hay envios
	envios := make([]*model.Envio, 0)

	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			return nil, err
		}

		envios = append(envios, &envio)
	}

	return envios, err
}

func (repository EnvioRepository) ObtenerEnvios(filtroEnvio *utils.FiltroEnvio) ([]*model.Envio, error) {
	//Desestructuramos el filtro
	patente := filtroEnvio.PatenteCamion
	estado := filtroEnvio.Estado
	ultimaParada := filtroEnvio.UltimaParada
	fechaCreacionDesde := filtroEnvio.FechaCreacionDesde
	fechaCreacionHasta := filtroEnvio.FechaCreacionHasta
	fechaUltimaActualizacionDesde := filtroEnvio.FechaUltimaActualizacionDesde
	fechaUltimaActualizacionHasta := filtroEnvio.FechaUltimaActualizacionHasta

	//Creamos el filtro para la base de datos
	filtro := bson.M{}

	//Solo filtra por patente si le pasamos un valor distinto de ""
	if patente != "" {
		filtro["patente_camion"] = patente
	}

	//Solo filtra por estado si le pasamos un estado no nulo
	if estado != "" {
		filtro["estado"] = estado
	}

	//Tomo la fecha de creacion en 0001-01-01 como la ausencia de filtro
	if !fechaCreacionDesde.IsZero() || !fechaCreacionHasta.IsZero() {
		filtroFecha := bson.M{}
		if !fechaCreacionDesde.IsZero() {
			filtroFecha["$gte"] = fechaCreacionDesde
		}
		if !fechaCreacionHasta.IsZero() {
			filtroFecha["$lte"] = fechaCreacionHasta
		}
		filtro["fecha_creacion"] = filtroFecha
	}

	//Si el estado es despachado, no tiene paradas y no tiene sentido filtrar por ultima parada
	if ultimaParada != "" && estado != model.ADespachar{
		customJavaScript := "this.paradas[this.paradas.length - 1].ciudad === '" + ultimaParada + "'"

		filtro["$where"] = customJavaScript
	}

	//Tomo la fecha de ultima actualizacion en 0001-01-01 como la ausencia de filtro
	if !fechaUltimaActualizacionDesde.IsZero() || !fechaUltimaActualizacionHasta.IsZero() {
		filtroFecha := bson.M{}
		if !fechaUltimaActualizacionDesde.IsZero() {
			filtroFecha["$gte"] = fechaUltimaActualizacionDesde
		}
		if !fechaUltimaActualizacionHasta.IsZero() {
			filtroFecha["$lte"] = fechaUltimaActualizacionHasta
		}
		filtro["fecha_ultima_actualizacion"] = filtroFecha
	}

	return repository.obtenerEnvios(filtro)
}

func (repository EnvioRepository) ObtenerEnvioPorId(envio *model.Envio) (*model.Envio, error) {
	filtro := bson.M{"_id": envio.ObjectId}

	envios, err := repository.obtenerEnvios(filtro)

	if err != nil {
		return &model.Envio{}, err
	}

	//Controlo que la lista este vacia
	if len(envios) == 0 {
		return &model.Envio{}, nil
	}

	return envios[0], nil
}

func (repository EnvioRepository) ObtenerCantidadEnviosPorEstado(estado model.EstadoEnvio) (int, error) {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")

	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)

	if err != nil {
		return 0, err
	}

	return int(cantidad), nil
}

func (repository EnvioRepository) ActualizarEnvio(envio *model.Envio) error {
	collection := repository.db.GetClient().Database("empresa").Collection("envios")
	filtro := bson.M{"_id": envio.ObjectId}

	//seteo la fecha de actualizacion
	envio.FechaUltimaActualizacion = time.Now()

	//Solo actualizamos ciertos campos
	actualizacion := bson.M{"$set": bson.M{
		"estado":                    envio.Estado,
		"fecha_ultima_actualizacion": envio.FechaUltimaActualizacion,
		"patente_camion":            envio.PatenteCamion,
		"pedidos":                   envio.Pedidos,
		"paradas":                   envio.Paradas,
	}}

	operacion, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el envio a actualizar")
	}

	return err
}
