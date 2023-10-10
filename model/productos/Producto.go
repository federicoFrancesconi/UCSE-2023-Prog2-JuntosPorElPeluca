package productos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto           int                `bson:"codigo_producto"`
	TipoDeProducto           TipoProducto       `bson:"tipo_producto"`
	Nombre                   string             `bson:"nombre"`
	PesoUnitario             float32            `bson:"peso_unitario"`
	PrecioUnitario           float32            `bson:"precio_unitario"`
	StockMinimo              int                `bson:"stock_minimo"`
	StockActual              int                `bson:"stock_actual"`
	FechaCreacion            time.Time          `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `bson:"fecha_ultima_actualizacion"`
	IdCreador                int                `bson:"id_creador"`
}
