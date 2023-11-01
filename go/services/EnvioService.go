package services

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"errors"
	"fmt"
	"time"
)

type EnvioServiceInterface interface {
	CrearEnvio(*dto.Envio) error
	ObtenerEnviosFiltrados(utils.FiltroEnvio) ([]*dto.Envio, error)
	ObtenerEnvioPorId(*dto.Envio) (*dto.Envio, error)
	ObtenerBeneficioEntreFechas(utils.FiltroEnvio) (float32, error)
	ObtenerCantidadEnviosPorEstado() ([]utils.CantidadEstado, error)
	AgregarParada(*dto.Envio) (bool, error)
	CambiarEstadoEnvio(*dto.Envio) (bool, error)
}

type EnvioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface, camionRepository repositories.CamionRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *EnvioService {
	return &EnvioService{
		envioRepository:    envioRepository,
		camionRepository:   camionRepository,
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
	}
}

func (service *EnvioService) CrearEnvio(envio *dto.Envio) error {
	envioCabeEnCamion, err := service.envioCabeEnCamion(envio)

	if err != nil {
		return err
	}

	if !envioCabeEnCamion {
		//Devuelve un error diciendo que el envio no cabe en el camion
		return errors.New("el envio no cabe en el camion")
	}

	//al crearlo coloco el envio en estado despachar
	envio.Estado = model.EstadoEnvio(model.ADespachar)

	//Cambio el estado de los pedidos del envio
	err = service.enviarPedidosDeEnvio(envio)

	if err != nil {
		return err
	}

	return service.envioRepository.CrearEnvio(envio.GetModel())
}

func (service *EnvioService) ObtenerEnviosFiltrados(filtroEnvio utils.FiltroEnvio) ([]*dto.Envio, error) {
	enviosDB, err := service.envioRepository.ObtenerEnviosFiltrados(filtroEnvio)

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

func (service *EnvioService) ObtenerEnvioPorId(envioConID *dto.Envio) (*dto.Envio, error) {
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envioConID.GetModel())

	var envio *dto.Envio

	if err != nil {
		return nil, err
	} else {
		envio = dto.NewEnvio(envioDB)
	}

	return envio, nil
}

func (service *EnvioService) envioCabeEnCamion(envio *dto.Envio) (bool, error) {
	//Primero buscamos el camion por patente
	camionConPatente := model.Camion{Patente: envio.PatenteCamion}
	camion, err := service.camionRepository.ObtenerCamionPorPatente(camionConPatente)

	if err != nil {
		return false, err
	}

	//Obtenemos el peso total de los pedidos
	var pesoTotal float32 = 0
	for _, idPedido := range envio.Pedidos {
		//Generamos el pedido para buscar
		pedidoParaBuscar := dto.Pedido{Id: idPedido}

		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoParaBuscar.GetModel())

		if err != nil {
			return false, err
		}

		//Calculo el peso del pedido sumando el peso de cada producto elegido
		var peso float32 = 0
		for _, producto := range pedido.ProductosElegidos {
			peso += producto.ObtenerPesoProductoPedido()
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

func (service *EnvioService) enviarPedidosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {
		err := service.enviarPedido(&dto.Pedido{Id: idPedido})
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *EnvioService) enviarPedido(pedidoPorEnviar *dto.Pedido) error {
	//Primero buscamos el pedido a enviar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorEnviar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Aceptado
	if pedido.Estado != model.Aceptado {
		return nil
	}

	//Cambia el estado del pedido a Para enviar, si es que no estaba ya en ese estado
	if pedido.Estado != model.ParaEnviar {
		pedido.Estado = model.ParaEnviar
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}

func (service *EnvioService) ObtenerBeneficioEntreFechas(filtro utils.FiltroEnvio) (float32, error) {
	//Pone por default las variables del filtro que no interesen
	filtro.PatenteCamion = ""
	filtro.UltimaParada = ""
	filtro.FechaCreacionDesde = time.Time{}
	filtro.FechaCreacionHasta = time.Time{}

	//Le agrega el estado despachado al filtro
	filtro.Estado = model.Despachado

	//Obtengo los envios despachados entre las dos fechas pasadas como parametro
	envios, err := service.ObtenerEnviosFiltrados(filtro)

	if err != nil {
		return 0, err
	}

	//Suma el precio de los pedidos de cada envio
	var beneficioBruto float32 = 0
	for _, envio := range envios {
		precioTotal, err := service.obtenerPrecioTotalProductosDeEnvio(envio)

		if err != nil {
			return 0, err
		}

		beneficioBruto += precioTotal
	}

	//Suma el costo de los envios
	var costoEnvios float32 = 0
	for _, envio := range envios {
		costoEnvio, err := service.obtenerCostoEnvio(envio)

		if err != nil {
			return 0, err
		}

		costoEnvios += costoEnvio
	}

	beneficioNeto := beneficioBruto - costoEnvios

	return beneficioNeto, nil
}

func (service *EnvioService) obtenerPrecioTotalProductosDeEnvio(envio *dto.Envio) (float32, error) {
	var precioTotal float32 = 0

	for _, idPedido := range envio.Pedidos {
		//Generamos el pedido para buscar
		pedidoParaBuscar := dto.Pedido{Id: idPedido}

		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoParaBuscar.GetModel())

		if err != nil {
			return 0, err
		}

		//Convierto el pedido a dto
		pedidoDTO := dto.NewPedido(pedido)

		precioTotal += pedidoDTO.ObtenerPecioTotal()
	}

	return precioTotal, nil
}

func (service *EnvioService) obtenerCostoEnvio(envio *dto.Envio) (float32, error) {
	//Obtiene el camion del envio para conocer el costoPorKilometro
	camion, err := service.camionRepository.ObtenerCamionPorPatente(model.Camion{Patente: envio.PatenteCamion})

	if err != nil {
		return 0, err
	}

	//Suma los kilometros de cada parada
	var kilometrosRecorridos int = 0
	for i := 0; i < len(envio.Paradas)-1; i++ {
		//Obtiene la distancia entre la parada i y la parada i+1
		kilometrosRecorridos += envio.Paradas[i].KmRecorridos
	}

	costoEnvio := camion.CostoPorKilometro * float32(kilometrosRecorridos)

	return costoEnvio, nil
}

func (service *EnvioService) ObtenerCantidadEnviosPorEstado() ([]utils.CantidadEstado, error) {
	//Por cada estado posible de envio, obtengo la cantidad de envios en ese estado
	cantidadEnviosADespachar, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.ADespachar)

	if err != nil {
		return nil, err
	}

	cantidadEnviosEnRuta, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.EnRuta)

	if err != nil {
		return nil, err
	}

	cantidadEnviosDespachados, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.Despachado)

	if err != nil {
		return nil, err
	}

	//Agrego los resultados a un array de CantidadEstado
	cantidadEnviosPorEstados := []utils.CantidadEstado{
		{Estado: "ADespachar", Cantidad: cantidadEnviosADespachar},
		{Estado: "EnRuta", Cantidad: cantidadEnviosEnRuta},
		{Estado: "Despachado", Cantidad: cantidadEnviosDespachados},
	}

	return cantidadEnviosPorEstados, nil
}

func (service *EnvioService) AgregarParada(envio *dto.Envio) (bool, error) {
	//En teoria, recibimos un envio que tiene solamente id y la nueva parada
	//Primero buscamos el envio por id
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envio.GetModel())

	if err != nil {
		return false, err
	}

	//Validamos que el envio esté en estado EnRuta
	if envioDB.Estado != model.EnRuta {
		return false, errors.New("el envio no esta en ruta")
	}

	//Agregamos la nueva parada al envio
	envioDB.Paradas = append(envioDB.Paradas, envio.Paradas[0].GetModel())

	//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
	return true, service.envioRepository.ActualizarEnvio(envioDB)
}

func (service *EnvioService) CambiarEstadoEnvio(envio *dto.Envio) (bool, error) {
	//El estado deseado es el que se pasa con el objeto envio como parametro
	estadoDeseado := envio.Estado

	//Buscamos el envio en la base de datos para conocer el estado real
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envio.GetModel())

	if err != nil {
		return false, err
	}

	//Si el estado del envio no es compatible con el deseado, devolvemos un error
	if (estadoDeseado == model.EnRuta && envioDB.Estado != model.ADespachar) || (estadoDeseado == model.Despachado && envioDB.Estado != model.EnRuta) {
		return false, errors.New("el envio no puede pasar al estado " + fmt.Sprint(estadoDeseado) + " si esta en estado " + fmt.Sprint(envioDB.Estado))
	}

	envioDB.Estado = estadoDeseado

	//Actualizamos el envio en la base de datos
	err = service.envioRepository.ActualizarEnvio(envioDB)

	if err != nil {
		return false, err
	}

	//Si el envio pasa a estado Despachado, finaliza el viaje, por lo que hay que hacer otras operaciones
	if estadoDeseado == model.Despachado {
		service.finalizarViaje(dto.NewEnvio(envioDB))
	}

	return true, nil
}

func (service *EnvioService) finalizarViaje(envio *dto.Envio) (bool, error) {
	//pasar pedidos a estado enviado
	err := service.entregarPedidosDeEnvio(envio)

	if err != nil {
		return false, err
	}

	//descontar stock de productos
	err = service.descontarStockProductosDeEnvio(envio)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (service *EnvioService) entregarPedidosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {

		//Descuenta el stock de los productos
		err := service.entregarPedido(&dto.Pedido{Id: idPedido})

		if err != nil {
			return err
		}
	}
	return nil
}

func (service *EnvioService) entregarPedido(pedidoPorEntregar *dto.Pedido) error {
	//Primero buscamos el pedido a entregar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorEntregar.GetModel())

	if err != nil {
		return err
	}

	//Valida que el pedido esté en estado Para enviar
	if pedido.Estado != model.ParaEnviar {
		return nil
	}

	//Cambia el estado del pedido a Enviado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Enviado {
		pedido.Estado = model.Enviado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}

func (service *EnvioService) descontarStockProductosDeEnvio(envio *dto.Envio) error {
	for _, idPedido := range envio.Pedidos {
		//Generamos el pedido para buscar
		pedidoParaBuscar := dto.Pedido{Id: idPedido}

		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoParaBuscar.GetModel())
		if err != nil {
			return err
		}

		for _, producto := range pedido.ProductosElegidos {
			err = service.descontarStockProducto(*dto.NewProductoPedido(&producto))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (service *EnvioService) descontarStockProducto(productoPedido dto.ProductoPedido) error {
	//Generamos un producto con el codigo del producto del pedido
	dtoProductoConId := dto.Producto{CodigoProducto: productoPedido.CodigoProducto}

	//Buscamos el producto del que hay que descontar la cantidad
	producto, err := service.productoRepository.ObtenerProductoPorCodigo(dtoProductoConId.GetModel())

	if err != nil {
		return err
	}

	//Modificamos el stock
	producto.StockActual = producto.StockActual - productoPedido.Cantidad

	//Actualizamos la base de datos
	return service.productoRepository.ActualizarProducto(*producto)
}
