package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
)

type CamionServiceInterface interface {
	CrearCamion(*dto.Camion, *dto.User) error
	ObtenerCamiones(utils.FiltroCamion) ([]*dto.Camion, error)
	ActualizarCamion(*dto.Camion) error
	EliminarCamion(*dto.Camion) error
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{camionRepository: camionRepository}
}

func (service *CamionService) CrearCamion(camion *dto.Camion, usuario *dto.User) error {
	return service.camionRepository.CrearCamion(camion.GetModel(), usuario.Codigo)
}

func (service *CamionService) ObtenerCamiones(filtro utils.FiltroCamion) ([]*dto.Camion, error) {
	camionesDB, err := service.camionRepository.ObtenerCamiones(filtro)

	if err != nil {
		return nil, err
	}

	//Inicializo la lista de camiones por si no hay ninguno
	camiones := make([]*dto.Camion, 0)

	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(camionDB)
		camiones = append(camiones, camion)
	}

	return camiones, nil
}

func (service *CamionService) ActualizarCamion(camion *dto.Camion) error {
	return service.camionRepository.ActualizarCamion(camion.GetModel())
}

func (service *CamionService) EliminarCamion(camionConPatente *dto.Camion) error {
	return service.camionRepository.EliminarCamion(camionConPatente.GetModel())
}
