package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model/envios"
)

type EnvioRepositoryInterface interface {
	CrearEnvio(envio *envios.Envio) error
	ObtenerEnvioPorId(id int) (*envios.Envio, error)
	ObtenerEnvios() ([]*envios.Envio, error)
	ActualizarEnvio(envio *envios.Envio) error
}

type EnvioRepository struct {
	db database.DB
}

func NewEnvioRepository(db database.DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}
