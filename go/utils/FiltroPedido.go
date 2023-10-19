package utils

import (
	"UCSE-2023-Prog2-TPIntegrador/model"
	"time"
)

type FiltroPedido struct {
	IdPedidos             []int
	Estado                model.EstadoPedido
	FechaCreacionComienzo time.Time
	FechaCreacionFin      time.Time
}
