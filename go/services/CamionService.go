package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/utils"
	"TPIntegrador/repositories"
	"errors"
)

type CamionServiceInterface interface {
	CrearCamion(*dto.Camion, *dto.User) error
	ObtenerCamiones(utils.FiltroCamion) ([]*dto.Camion, error)
	ActualizarCamion(*dto.Camion, *dto.User) error
	EliminarCamion(*dto.Camion, *dto.User) error
}

type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{camionRepository: camionRepository}
}

func (service *CamionService) CrearCamion(camion *dto.Camion, usuario *dto.User) error {
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para crear un camion")
	}

	//Le agregamos el codigo del usuario que lo creo
	camion.IdCreador = usuario.Codigo

	return service.camionRepository.CrearCamion(camion.GetModel())
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

func (service *CamionService) ActualizarCamion(camion *dto.Camion, usuario *dto.User) error {
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para actualizar un camion")
	}

	return service.camionRepository.ActualizarCamion(camion.GetModel())
}

func (service *CamionService) EliminarCamion(camionConPatente *dto.Camion, usuario *dto.User) error {
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para eliminar un camion")
	}

	return service.camionRepository.EliminarCamion(camionConPatente.GetModel())
}

func (service *CamionService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == "Administrador"
}
