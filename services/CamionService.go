package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type CamionInterface interface {
	ObtenerCamiones() []*dto.Camion
	ObtenerCamionPorPatente(patente string) *dto.Camion
	CrearCamion(camion *dto.Camion) bool
	ActualizarCamion(camion *dto.Camion) bool
	EliminarCamion(patente string) bool
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{camionRepository: camionRepository}
}

func (service *CamionService) ObtenerCamiones() []*dto.Camion {
	//Falta controlar el error
	camionesDB, _ := service.camionRepository.ObtenerCamiones()
	var camiones []*dto.Camion
	for _, camionDB := range camionesDB {
		camion := dto.NewCamion(camionDB)
		camiones = append(camiones, camion)
	}
	return camiones
}

func (service *CamionService) ObtenerCamionPorPatente(patente string) *dto.Camion {
	camionDB, err := service.camionRepository.ObtenerCamionPorPatente(patente)
	var camion *dto.Camion
	if err == nil {
		camion = dto.NewCamion(camionDB)
	}
	return camion
}

func (service *CamionService) CrearCamion(camion *dto.Camion) bool {
	service.camionRepository.CrearCamion(camion.GetModel())
	return true
}

func (service *CamionService) ActualizarCamion(camion *dto.Camion) bool {
	service.camionRepository.ActualizarCamion(camion.GetModel())
	return true
}

func (service *CamionService) EliminarCamion(patente string) bool {
	service.camionRepository.EliminarCamion(patente)
	return true
}
