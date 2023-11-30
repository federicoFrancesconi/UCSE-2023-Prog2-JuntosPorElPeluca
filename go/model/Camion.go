package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	Patente                  string             `bson:"patente"`
	PesoMaximo               int                `bson:"peso_maximo"`
	CostoPorKilometro        float64            `bson:"costo_por_kilometro"`
	FechaCreacion            time.Time          `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `bson:"fecha_ultima_actualizacion"`
	IdCreador                string             `bson:"id_creador"`
	EstaActivo               bool               `bson:"esta_activo"`
}

//TODO: funcion para validar,
