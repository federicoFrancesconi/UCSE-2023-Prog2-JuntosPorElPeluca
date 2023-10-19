package utils

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type FiltroPedido struct {
	IdPedidos             []int
	IdEnvio               int
	Estado                model.EstadoPedido
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
