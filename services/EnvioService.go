package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
)

type EnvioInterface interface {
	ObtenerEnvios() ([]*dto.Envio, error)
	ObtenerEnvioPorId(id int) (*dto.Envio, error)
	CrearEnvio(envio *dto.Envio) error
	AgregarParada(envio *dto.Envio) (bool, error)
	IniciarViaje(envio *dto.Envio) (bool, error)
	FinalizarViaje(envio *dto.Envio) (bool, error)
}

type EnvioService struct {
	envioRepository repositories.EnvioRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *EnvioService {
	return &EnvioService{envioRepository: envioRepository}
}

func (service *EnvioService) ObtenerEnvios() ([]*dto.Envio, error) {
	enviosDB, err := service.envioRepository.ObtenerEnvios()

	if err != nil {
		return nil, err
	}

	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios, nil
}

func (service *EnvioService) ObtenerEnvioPorId(id int) (*dto.Envio, error) {
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(id)
	var envio *dto.Envio
	if err != nil {
		return nil, err
	} else {
		envio = dto.NewEnvio(envioDB)
	}
	return envio, nil
}

func (service *EnvioService) CrearEnvio(envio *dto.Envio) error {
	for _, pedido := range envio.Pedidos {
		//pasar el pedido en estado para enviar

		//verificar que la suma de los pesos de cada producto en pedido no sobrepase el limite de peso del vehiculo
		var sumaPesos float32 = 0
		for _, producto := range pedido.ProductosElegidos {
			sumaPesos += producto.PesoUnitario
		}

		//ver como obtener el camion para verificar la suma de los pesos
	}

	//al crearlo coloco el envio en estado despachar
	envio.Estado = model.EstadoEnvio(model.ParaEnviar)

	return service.envioRepository.CrearEnvio(envio.GetModel())
}

func (service *EnvioService) AgregarParada(envio *dto.Envio) (bool, error) {
	if envio.Estado != model.EnRuta {
		return false, nil
	}

	return true, service.envioRepository.ActualizarEnvio(envio.GetModel())
}

func (service *EnvioService) IniciarViaje(envio *dto.Envio) (bool, error) {
	if envio.Estado != model.ADespachar {
		return false, nil
	}

	envio.Estado = model.EnRuta

	return true, service.envioRepository.ActualizarEnvio(envio.GetModel())
}

func (service *EnvioService) FinalizarViaje(envio *dto.Envio) (bool, error) {
	if envio.Estado == model.Despachado {
		return false, nil
	}

	envio.Estado = model.Despachado

	service.envioRepository.ActualizarEnvio(envio.GetModel())
	//pasar pedidos a estado enviado
	//descontar stock de productos

	return true, nil
}
