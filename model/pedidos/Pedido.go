package pedidos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"UCSE-2023-Prog2-TPIntegrador/model/productos"
)

type Pedido struct {
	ObjectId                 primitive.ObjectID         `bson:"_id,omitempty"`
	FechaCreacion            time.Time                  `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time                  `bson:"fecha_ultima_actualizacion"`
	IdCreador                int                        `bson:"id_creador"`
	ProductosElegidos        []productos.ProductoPedido `bson:"productos_elegidos"`
	CiudadDestino            string                     `bson:"ciudad_destino"`
	Estado                   EstadoPedido               `bson:"estado"`
}
