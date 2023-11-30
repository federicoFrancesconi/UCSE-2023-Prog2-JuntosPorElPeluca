package utils

import (
	"TPIntegrador/model"
	"time"
)

type FiltroPedido struct {
	IdPedidos             []string
	IdEnvio               string
	Estado                model.EstadoPedido
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
