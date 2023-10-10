package envios

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/model/pedidos"
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
	Pedidos                  []pedidos.Pedido   `bson:"pedidos"`
	IdCreador                int                `bson:"idCreador"`
	Estado                   EstadoEnvio        `bson:"estado"`
}
