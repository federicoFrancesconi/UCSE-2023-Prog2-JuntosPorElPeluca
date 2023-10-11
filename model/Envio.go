package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	//Por qué el Id es un string?
	Id                       string             `bson:"id"`
	FechaCreacion            time.Time          `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `bson:"fecha_ultima_actualizacion"`
	PatenteCamion            string             `bson:"patente_camion"`
	Paradas                  []Parada     `bson:"paradas"`
	Pedidos                  []Pedido   `bson:"pedidos"`
	IdCreador                int                `bson:"id_creador"`
	Estado                   EstadoEnvio        `bson:"estado"`
}
