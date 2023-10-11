package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type EnvioInterface interface {
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorId(id string) *dto.Envio
	CrearEnvio(envio *dto.Envio) bool
	ActualizarEnvio(envio *dto.Envio) bool
	EliminarEnvio(id string) bool
}

type EnvioService struct {
	envioRepository repositories.EnvioRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *EnvioService {
	return &EnvioService{envioRepository: envioRepository}
}

func (service *EnvioService) ObtenerEnvios() []*dto.Envio {
	enviosDB, _ := service.envioRepository.ObtenerEnvios()
	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios
}

func (service *EnvioService) ObtenerEnvioPorId(id string) *dto.Envio {
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(id)
	var envio *dto.Envio
	if err == nil {
		envio = dto.NewEnvio(envioDB)
	}
	return envio
}

func (service *EnvioService) CrearEnvio(envio *dto.Envio) bool {
	service.envioRepository.CrearEnvio(envio.GetModel())
	return true
}

func (service *EnvioService) ActualizarEnvio(envio *dto.Envio) bool {
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}

func (service *EnvioService) EliminarEnvio(id string) bool {
	service.envioRepository.EliminarEnvio(id)
	return true
}
