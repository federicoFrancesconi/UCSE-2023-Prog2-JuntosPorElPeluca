package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	ObjectId                 primitive.ObjectID         `bson:"_id,omitempty"`
	FechaCreacion            time.Time                  `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time                  `bson:"fecha_ultima_actualizacion"`
	IdCreador                int                        `bson:"id_creador"`
	ProductosElegidos        []ProductoPedido `bson:"productos_elegidos"`
	CiudadDestino            string                     `bson:"ciudad_destino"`
	Estado                   EstadoPedido               `bson:"estado"`
}
