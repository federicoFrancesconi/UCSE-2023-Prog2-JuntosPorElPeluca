package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ObjectId                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CodigoProducto           int                `bson:"codigoProducto" json:"codigoProducto"`
	TipoDeProducto           TipoProducto       `bson:"tipoDeProducto" json:"tipoDeProducto"`
	Nombre                   string             `bson:"nombre" json:"nombre"`
	PesoUnitario             int                `bson:"pesoUnitario" json:"pesoUnitario"`
	PrecioUnitario           float32            `bson:"precioUnitario" json:"precioUnitario"`
	StockMinimo              int                `bson:"stockMinimo" json:"stockMinimo"`
	FechaCreacion            time.Time          `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaUltimaActualizacion time.Time          `bson:"fechaUltimaActualizacion" json:"fechaUltimaActualizacion"`
	IdCreador                int                `bson:"idCreador" json:"idCreador"`
}
