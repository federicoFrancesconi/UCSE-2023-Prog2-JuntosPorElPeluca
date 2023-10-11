package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type CamionInterface interface {
	ObtenerCamiones() ([]*dto.Camion, error)
	ObtenerCamionPorPatente(patente string) (*dto.Camion, error)
	CrearCamion(camion *dto.Camion) error
	ActualizarCamion(camion *dto.Camion) error
	EliminarCamion(patente string) error
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{camionRepository: camionRepository}
}

func (service *CamionService) ObtenerCamiones() ([]*dto.Camion, error) {
	//Falta controlar el error
	camionesDB, err := service.camionRepository.ObtenerCamiones()

	if err != nil {
		return nil, err
	}

	var camiones []*dto.Camion
	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(camionDB)
		camiones = append(camiones, camion)
	}
	return camiones, nil
}

func (service *CamionService) ObtenerCamionPorPatente(patente string) (*dto.Camion, error) {
	camionDB, err := service.camionRepository.ObtenerCamionPorPatente(patente)
	var camion *dto.Camion
	if err != nil {
		return nil, err
	}

	camion = dto.NewCamion(camionDB)

	return camion, nil
}

func (service *CamionService) CrearCamion(camion *dto.Camion) error {
	return service.camionRepository.CrearCamion(camion.GetModel())
}

func (service *CamionService) ActualizarCamion(camion *dto.Camion) error {
	return service.camionRepository.ActualizarCamion(camion.GetModel())
}

func (service *CamionService) EliminarCamion(patente string) error {
	return service.camionRepository.EliminarCamion(patente)
}
