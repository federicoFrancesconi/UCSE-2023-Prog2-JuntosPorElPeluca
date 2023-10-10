package envios

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/model/pedidos"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	Id                       string             `bson:"id"`
	FechaCreacion            time.Time          `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `bson:"fecha_ultima_actualizacion"`
	PatenteCamion            string             `bson:"patente_camion"`
	Paradas                  []model.Parada     `bson:"paradas"`
	Pedidos                  []pedidos.Pedido   `bson:"pedidos"`
	IdCreador                int                `bson:"id_creador"`
	Estado                   EstadoEnvio        `bson:"estado"`
}
