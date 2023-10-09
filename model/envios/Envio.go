package envios

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	Id                       int                `bson:"id"`
	FechaCreacion            time.Time          `bson:"fechaCreacion"`
	FechaUltimaActualizacion time.Time          `bson:"fechaUltimaActualizacion"`
	PatenteCamion            string             `bson:"patente-amion"`
	Paradas                  []model.Parada     `bson:"paradas"`
	IdCreador                int                `bson:"idCreador"`
	Estado                   EstadoEnvio        `bson:"estado"`
}
