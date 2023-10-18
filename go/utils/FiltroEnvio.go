package utils

import (
	"time"
	"UCSE-2023-Prog2-TPIntegrador/model"
)

type FiltroEnvio struct {
	PatenteCamion            string
	Estado                   model.EstadoEnvio
	UltimaParada             string
	FechaCreacionComienzo    time.Time
	FechaCreacionFin         time.Time
}