package pedidos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"UCSE-2023-Prog2-TPIntegrador/model/productos"
)

type Pedido struct {
	ObjectId                 primitive.ObjectID         `bson:"_id,omitempty"`
	FechaCreacion            time.Time                  `bson:"fechaCreacion"`
	FechaUltimaActualizacion time.Time                  `bson:"fechaUltimaActualizacion"`
	IdCreador                int                        `bson:"idCreador"`
	ProductosElegidos        []productos.ProductoPedido `bson:"productosPedido"`
	CiudadDestino            string                     `bson:"ciudadDestino"`
	Estado                   EstadoPedido               `bson:"estado"`
}
