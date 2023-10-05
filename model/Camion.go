package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	Patente                  string             `bson:"patente"`
	PesoMaximo               int                `bson:"pesoMaximo"`
	FechaCreacion            time.Time          `bson:"fechaCreacion"`
	FechaUltimaActualizacion time.Time          `bson:"fechaUltimaActualizacion"`
	IdCreador                int                `bson:"idCreador"`
}
