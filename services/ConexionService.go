package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
)

type ConexionServiceInterface interface {
	EnvioCabeEnCamion(*dto.Envio) (bool, error)
	EntregarPedidosDeEnvio(*dto.Envio) error
	DescontarStockProductosDeEnvio(*dto.Envio) error
}

type ConexionService struct {
	camionService CamionServiceInterface
	pedidoService PedidoServiceInterface
	productoService ProductoServiceInterface
}

func NewConexionService(camionService CamionServiceInterface, pedidoService PedidoServiceInterface, productoService ProductoServiceInterface) *ConexionService {
	return &ConexionService{
		camionService: camionService,
		pedidoService: pedidoService,
		productoService: productoService,
	}
}

func (service *ConexionService) EnvioCabeEnCamion(envio *dto.Envio) (bool, error) {
	//Primero buscamos el camion por patente
	camion, err := service.camionService.ObtenerCamionPorPatente(envio.PatenteCamion)
	if err != nil {
		return false, err
	}

	//Obtenemos el peso total de los pedidos
	var pesoTotal float32 = 0
	for _, idPedido := range envio.Pedidos {
		peso, err := service.pedidoService.ObtenerPesoPedido(idPedido)

		if err != nil {
			return false, err
		}

		pesoTotal += peso
	}

	//Verificamos si el peso total de los pedidos es menor o igual al peso maximo del camion
	if pesoTotal <= float32(camion.PesoMaximo) {
		return true, nil
	} else {
		return false, nil
	}
}

func (service *ConexionService) EntregarPedidosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {
		err := service.pedidoService.EntregarPedido(idPedido)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *ConexionService) DescontarStockProductosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {
		pedido, err := service.pedidoService.ObtenerPedidoPorId(idPedido)
		if err != nil {
			return err
		}

		for _, producto := range pedido.ProductosElegidos {
			err = service.productoService.DescontarStockProducto(producto.CodigoProducto, producto.Cantidad)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
