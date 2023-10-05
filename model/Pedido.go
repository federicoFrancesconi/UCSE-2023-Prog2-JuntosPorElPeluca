package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	FechaCreacion            time.Time          `bson:"fechaCreacion"`
	FechaUltimaActualizacion time.Time          `bson:"fechaUltimaActualizacion"`
	IdCreador                int                `bson:"idCreador"`
	ProductosPedido          []ProductoPedido   `bson:"productosPedido"`
	CiudadDestino            string             `bson:"ciudadDestino"`
	Estado                   string             `bson:"estado"`
}
