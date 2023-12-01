package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/model"
	"TPIntegrador/repositories"
	"TPIntegrador/utils"
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
	envioRepository repositories.EnvioRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface, envioRepository repositories.EnvioRepositoryInterface) *CamionService {
	return &CamionService{
		camionRepository: camionRepository,
		envioRepository: envioRepository,
	}
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
	//Aseguramos que el filtro tenga el campo esta_activo en true
	filtro.EstaActivo = true
	filtro.FiltrarPorEstaActivo = true

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

//En lugar de eliminar el camion, actualiza el campo esta_activo a false
func (service *CamionService) EliminarCamion(camionConPatente *dto.Camion, usuario *dto.User) error {
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para eliminar un camion")
	}

	//Creamos el filtro para obtener el camion
	filtro := utils.FiltroCamion{Patente: camionConPatente.Patente}

	//Obtengo el camion de la base de datos
	camiones, err := service.camionRepository.ObtenerCamiones(filtro)

	if err != nil {
		return err
	}

	//Si no existe el camion, devuelvo un error
	if len(camiones) == 0 {
		return errors.New("no existe el camion")
	}

	camion := camiones[0]

	//Valido que el camion no tenga envios actualmente
	err, tieneEnvios := service.camionTieneEnviosActualmente(dto.NewCamion(camion))

	if err != nil {
		return err
	}

	//Si tiene envios, devuelvo un error
	if tieneEnvios {
		return errors.New("el camion tiene envios actualmente")
	}

	//Actualizo el campo esta_activo a false
	camion.EstaActivo = false

	//Actualizo el camion
	return service.camionRepository.ActualizarCamion(camion)
}

func (service *CamionService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == "Administrador"
}

func (service *CamionService) camionTieneEnviosActualmente(camion *dto.Camion) (error, bool) {
	//Creamos el filtro para obtener los envios
	filtro := utils.FiltroEnvio{PatenteCamion: camion.Patente, Estado: model.ADespachar}

	//Obtengo los envios de la base de datos
	enviosADespachar, err := service.envioRepository.ObtenerEnviosFiltrados(filtro)

	if err != nil {
		return errors.New("error al obtener los envios a despachar"), false
	}

	//Hacemos lo mismo para los envios que estan En Ruta
	filtro.Estado = model.EnRuta

	enviosEnRuta, err := service.envioRepository.ObtenerEnviosFiltrados(filtro)

	if err != nil {
		return errors.New("error al obtener los envios en ruta"), false
	}

	//Si no hay envios, devuelvo false
	if len(enviosADespachar) == 0 && len(enviosEnRuta) == 0 {
		return nil, false
	}

	return nil, true
}
