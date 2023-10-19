package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type CamionServiceInterface interface {
	CrearCamion(*dto.Camion) error
	ObtenerCamiones() ([]*dto.Camion, error)
	ObtenerCamionPorPatente(*dto.Camion) (*dto.Camion, error)
	ActualizarCamion(*dto.Camion) error
	EliminarCamion(*dto.Camion) error
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{camionRepository: camionRepository}
}

func (service *CamionService) CrearCamion(camion *dto.Camion) error {
	return service.camionRepository.CrearCamion(camion.GetModel())
}

func (service *CamionService) ObtenerCamiones() ([]*dto.Camion, error) {
	camionesDB, err := service.camionRepository.ObtenerTodosLosCamiones()

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

func (service *CamionService) ObtenerCamionPorPatente(camionConPatente *dto.Camion) (*dto.Camion, error) {
	camionDB, err := service.camionRepository.ObtenerCamionPorPatente(camionConPatente.GetModel())
	
	var camion *dto.Camion
	if err != nil {
		return nil, err
	}

	camion = dto.NewCamion(camionDB)

	return camion, nil
}

func (service *CamionService) ActualizarCamion(camion *dto.Camion) error {
	return service.camionRepository.ActualizarCamion(camion.GetModel())
}

func (service *CamionService) EliminarCamion(camionConPatente *dto.Camion) error {
	return service.camionRepository.EliminarCamion(camionConPatente.GetModel())
}
