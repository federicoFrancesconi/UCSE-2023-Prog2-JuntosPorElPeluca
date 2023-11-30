package dto

import (
	"TPIntegrador/model"
	"TPIntegrador/utils"
	"time"
)

type Producto struct {
	CodigoProducto           string             `json:"codigo_producto"`
	TipoDeProducto           model.TipoProducto `json:"tipo_producto"`
	Nombre                   string             `json:"nombre"`
	PesoUnitario             float32            `json:"peso_unitario"`
	PrecioUnitario           float32            `json:"precio_unitario"`
	StockMinimo              int                `json:"stock_minimo"`
	StockActual              int                `json:"stock_actual"`
	FechaCreacion            time.Time          `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `json:"fecha_ultima_actualizacion"`
	IdCreador                string             `json:"id_creador"`
}

// Crea el dto a partir del modelo
func NewProducto(producto *model.Producto) *Producto {
	return &Producto{
		CodigoProducto:           utils.GetStringIDFromObjectID(producto.ObjectId),
		TipoDeProducto:           producto.TipoDeProducto,
		Nombre:                   producto.Nombre,
		PrecioUnitario:           producto.PrecioUnitario,
		PesoUnitario:             producto.PesoUnitario,
		StockMinimo:              producto.StockMinimo,
		StockActual:              producto.StockActual,
		FechaCreacion:            producto.FechaCreacion,
		FechaUltimaActualizacion: producto.FechaUltimaActualizacion,
		IdCreador:                producto.IdCreador,
	}
}

// Crea el modelo a partir del dto
func (producto Producto) GetModel() model.Producto {
	return model.Producto{
		ObjectId:                 utils.GetObjectIDFromStringID(producto.CodigoProducto),
		TipoDeProducto:           producto.TipoDeProducto,
		Nombre:                   producto.Nombre,
		PrecioUnitario:           producto.PrecioUnitario,
		PesoUnitario:             producto.PesoUnitario,
		StockMinimo:              producto.StockMinimo,
		StockActual:              producto.StockActual,
		FechaCreacion:            producto.FechaCreacion,
		FechaUltimaActualizacion: producto.FechaUltimaActualizacion,
		IdCreador:                producto.IdCreador,
	}
}
