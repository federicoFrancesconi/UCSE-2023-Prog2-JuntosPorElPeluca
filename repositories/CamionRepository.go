package repositories

import (
	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/model"
)

type CamionRepositoryInterface interface {
	CrearCamion(camion *model.Camion) error
	ObtenerCamionPorPatente(patente string) (*model.Camion, error)
	ObtenerCamiones() ([]*model.Camion, error)
	ActualizarCamion(camion *model.Camion) error
}

type CamionRepository struct {
	db database.DB
}

func NewCamionRepository(db database.DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}
