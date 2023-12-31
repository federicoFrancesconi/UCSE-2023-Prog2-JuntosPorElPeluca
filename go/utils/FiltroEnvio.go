package utils

import (
	"TPIntegrador/model"
	"time"
)

type FiltroEnvio struct {
	PatenteCamion                 string
	Estado                        model.EstadoEnvio
	UltimaParada                  string
	FechaCreacionDesde            time.Time
	FechaCreacionHasta            time.Time
	FechaUltimaActualizacionDesde time.Time
	FechaUltimaActualizacionHasta time.Time
}
