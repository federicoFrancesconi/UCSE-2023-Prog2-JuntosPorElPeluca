package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	Id                       int                `bson:"id"`
	FechaCreacion            time.Time          `bson:"fechaCreacion"`
	FechaUltimaActualizacion time.Time          `bson:"fechaUltimaActualizacion"`
	PatenteCamion            string             `bson:"patente-amion"`
	Paradas                  []Parada           `bson:"paradas"`
	IdCreador                int                `bson:"idCreador"`
	Estado                   EstadoEnvio        `bson:"estado"`
}
