package services

import (
	"TPIntegrador/dto"
	"TPIntegrador/model"
	"TPIntegrador/repositories"
	"TPIntegrador/utils"
	"errors"
	"fmt"
	"time"
)

type EnvioServiceInterface interface {
	CrearEnvio(*dto.Envio, *dto.User) error
	ObtenerEnvios(utils.FiltroEnvio) ([]*dto.Envio, error)
	ObtenerEnvioPorId(*dto.Envio) (*dto.Envio, error)
	ObtenerBeneficioTemporal(utils.FiltroEnvio) (dto.BeneficioTemporal, error)
	ObtenerCantidadEnviosPorEstado() ([]utils.CantidadEstado, error)
	AgregarParada(*dto.NuevaParada, *dto.User) (bool, error)
	CambiarEstadoEnvio(*dto.Envio, *dto.User) (bool, error)
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

func (service *EnvioService) CrearEnvio(envio *dto.Envio, usuario *dto.User) error {
	//valido que el envio lo este creando un camionero
	if !service.validarRol(usuario) {
		return errors.New("el usuario no tiene permisos para crear un envio")
	}

	envioCabeEnCamion, err := service.envioCabeEnCamion(envio)

	if err != nil {
		return err
	}

	if !envioCabeEnCamion {
		//Devuelve un error diciendo que el envio no cabe en el camion
		return errors.New("el envio no cabe en el camion")
	}

	//al crearlo coloco el envio en estado despachar
	envio.Estado = model.ADespachar

	//Indicamos el usuario que creo el envio
	envio.IdCreador = usuario.Codigo

	//Cambio el estado de los pedidos del envio
	err = service.enviarPedidosDeEnvio(envio)

	if err != nil {
		return err
	}

	//descontar stock de productos
	err = service.descontarStockProductosDeEnvio(envio)

	if err != nil {
		return err
	}

	return service.envioRepository.CrearEnvio(envio.GetModel())
}

func (service *EnvioService) ObtenerEnvios(filtroEnvio utils.FiltroEnvio) ([]*dto.Envio, error) {
	//Validamos el estado que se paso para filtrar
	if filtroEnvio.Estado != "" {
		if !model.EsUnEstadoEnvioValido(filtroEnvio.Estado) {
			return nil, errors.New("el estado ingresado para filtrar no es válido")
		}
	}

	fechaDesde := filtroEnvio.FechaCreacionDesde
	fechaHasta := filtroEnvio.FechaCreacionHasta

	//Valida que la fecha desde sea menor a la fecha hasta, solo si se pasaron ambas para filtrar
	if fechaDesde.Year() != 1 && fechaHasta.Year() != 1 && fechaDesde.After(fechaHasta) {
		return nil, errors.New("la fecha desde debe ser menor o igual a la fecha hasta")
	}

	enviosDB, err := service.envioRepository.ObtenerEnvios(&filtroEnvio)

	if err != nil {
		return nil, err
	}

	//Inicializamos el array de envios por si no hay ninguno
	envios := []*dto.Envio{}

	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(*envioDB)
		envios = append(envios, envio)
	}
	
	return envios, nil
}

func (service *EnvioService) ObtenerEnvioPorId(envioConID *dto.Envio) (*dto.Envio, error) {
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envioConID.GetModel())

	//Inicializamos el envio por si no hay ninguno
	var envio *dto.Envio = &dto.Envio{}

	if err != nil {
		return nil, err
	} else {
		envio = dto.NewEnvio(*envioDB)
	}

	return envio, nil
}

func (service *EnvioService) envioCabeEnCamion(envio *dto.Envio) (bool, error) {
	//Primero buscamos el camion por patente
	filtroPorPatente := utils.FiltroCamion{Patente: envio.PatenteCamion, EstaActivo: true, FiltrarPorEstaActivo: true}

	camiones, err := service.camionRepository.ObtenerCamiones(filtroPorPatente)

	if err != nil {
		return false, err
	}

	//Si no existe el camion, devolvemos un error
	if len(camiones) == 0 {
		return false, errors.New("no existe el camion, o bien ha sido dado de baja")
	}

	camion := camiones[0]

	//Obtenemos el peso total de los pedidos
	var pesoTotal float64 = 0
	for _, idPedido := range envio.Pedidos {
		//Generamos el pedido para buscar
		pedidoParaBuscar := dto.Pedido{Id: idPedido}

		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoParaBuscar.GetModel())

		if err != nil {
			return false, err
		}

		//Calculo el peso del pedido sumando el peso de cada producto elegido
		var peso float64 = 0
		for _, producto := range pedido.ProductosElegidos {
			peso += producto.ObtenerPesoProductoPedido()
		}

		pesoTotal += peso
	}

	//Verificamos si el peso total de los pedidos es menor o igual al peso maximo del camion
	if pesoTotal <= float64(camion.PesoMaximo) {
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
		return errors.New("error buscando el pedido en la DB: " + err.Error())
	}

	//Valida que el pedido esté en estado Aceptado
	if pedido.Estado != model.Aceptado {
		return errors.New("el pedido " + pedidoPorEnviar.Id + " no se encuentra en estado Aceptado")
	}

	//Cambia el estado del pedido a Para enviar, si es que no estaba ya en ese estado
	if pedido.Estado != model.ParaEnviar {
		pedido.Estado = model.ParaEnviar
	}

	//Actualiza el pedido en la base de datos
	err = service.pedidoRepository.ActualizarPedido(pedido)

	if err != nil {
		return errors.New("error actualizando el pedido en la DB: " + err.Error())
	}

	return nil
}

func (service *EnvioService) ObtenerBeneficioTemporal(filtro utils.FiltroEnvio) (dto.BeneficioTemporal, error) {
	//Inicializamos el beneficio temporal
	beneficioTemporal := dto.BeneficioTemporal{}
	fechaDesde := filtro.FechaUltimaActualizacionDesde
	fechaHasta := filtro.FechaUltimaActualizacionHasta

	//Valida que la fecha desde sea menor a la fecha hasta
	if fechaDesde.After(fechaHasta) {
		return beneficioTemporal, errors.New("la fecha desde debe ser menor o igual a la fecha hasta")
	}

	//Obtiene el beneficio anual
	//Por cada año en el rango de fechas, obtiene el beneficio
	for año := fechaDesde.Year(); año <= fechaHasta.Year(); año++ {
		filtroPorAño := utils.FiltroEnvio{
			FechaUltimaActualizacionDesde: time.Date(año, 1, 1, 0, 0, 0, 0, time.UTC),
			FechaUltimaActualizacionHasta: time.Date(año, 12, 31, 23, 59, 59, 999999999, time.UTC),
		}

		montoBeneficioAnual, err := service.obtenerBeneficioEntreFechas(filtroPorAño)

		if err != nil {
			return beneficioTemporal, err
		}

		beneficioAnual := dto.BeneficioAnual{Año: año, Monto: montoBeneficioAnual}

		beneficioTemporal.BeneficiosAnuales = append(beneficioTemporal.BeneficiosAnuales, beneficioAnual)
	}

	//Hacemos lo mismo con los meses
	for año, mes := fechaDesde.Year(), fechaDesde.Month(); !((año == fechaHasta.Year()) && (mes > fechaHasta.Month()) || año > fechaHasta.Year()); {
		filtroPorMes := utils.FiltroEnvio{
			FechaUltimaActualizacionDesde: time.Date(año, mes, 1, 0, 0, 0, 0, time.UTC),
			FechaUltimaActualizacionHasta: time.Date(año, mes+1, 0, 23, 59, 59, 999999999, time.UTC),
		}

		montoBeneficioMensual, err := service.obtenerBeneficioEntreFechas(filtroPorMes)

		if err != nil {
			return beneficioTemporal, err
		}

		beneficioMensual := dto.BeneficioMensual{Mes: int(mes), Monto: montoBeneficioMensual}

		beneficioTemporal.BeneficiosMensuales = append(beneficioTemporal.BeneficiosMensuales, beneficioMensual)

		//Pasamos al mes siguiente
		if mes == 12 {
			año++
			mes = 1
		} else {
			mes++
		}
	}

	return beneficioTemporal, nil
}

func (service *EnvioService) obtenerBeneficioEntreFechas(filtro utils.FiltroEnvio) (float64, error) {
	//Le agrega el estado despachado al filtro, ya que el beneficio lo tienen los despachados
	filtro.Estado = model.Despachado

	//Obtengo los envios despachados entre las dos fechas pasadas como parametro
	envios, err := service.ObtenerEnvios(filtro)

	if err != nil {
		return 0, err
	}

	//Valido que haya envios despachados
	// if len(envios) == 0 {
	// 	return 0, errors.New("no hay envios despachados entre las fechas ingresadas")
	// }

	//Suma el precio de los pedidos de cada envio
	var beneficioBruto float64 = 0
	for _, envio := range envios {
		precioTotal, err := service.obtenerPrecioTotalProductosDeEnvio(envio)

		if err != nil {
			return 0, err
		}

		beneficioBruto += precioTotal
	}

	//Suma el costo de los envios
	var costoEnvios float64 = 0
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

func (service *EnvioService) obtenerPrecioTotalProductosDeEnvio(envio *dto.Envio) (float64, error) {
	var precioTotal float64 = 0

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

func (service *EnvioService) obtenerCostoEnvio(envio *dto.Envio) (float64, error) {
	//Obtiene el camion del envio para conocer el costoPorKilometro
	//Para el costo no es necesario filtrar por activo, ya que el camion puede estar dado de baja y el envio ya creado
	filtroPorPatente := utils.FiltroCamion{Patente: envio.PatenteCamion, FiltrarPorEstaActivo: false}
	camiones, err := service.camionRepository.ObtenerCamiones(filtroPorPatente)

	if err != nil {
		return 0, err
	}

	//Si no existe el camion, devolvemos un error
	if len(camiones) == 0 {
		return 0, errors.New("no existe el camion")
	}

	camion := camiones[0]

	//Suma los kilometros de cada parada
	var kilometrosRecorridos int = 0
	for i := 0; i < len(envio.Paradas)-1; i++ {
		//Obtiene la distancia entre la parada i y la parada i+1
		kilometrosRecorridos += envio.Paradas[i].KmRecorridos
	}

	costoEnvio := camion.CostoPorKilometro * float64(kilometrosRecorridos)

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
		{Estado: string(model.ADespachar), Cantidad: cantidadEnviosADespachar},
		{Estado: string(model.EnRuta), Cantidad: cantidadEnviosEnRuta},
		{Estado: string(model.Despachado), Cantidad: cantidadEnviosDespachados},
	}

	return cantidadEnviosPorEstados, nil
}

func (service *EnvioService) AgregarParada(parada *dto.NuevaParada, usuario *dto.User) (bool, error) {
	//Recibimos la parada con el id del envioSoloId a ingresarla
	envioSoloId := dto.Envio{Id: parada.IdEnvio}

	//Primero buscamos el envio por id
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envioSoloId.GetModel())

	if err != nil {
		return false, err
	}

	//Validamos el rol del usuario
	if !service.validarRol(usuario) {
		return false, errors.New("el usuario no tiene permisos para agregar una parada")
	}

	//Validamos que el envio esté en estado EnRuta
	if envioDB.Estado != model.EnRuta {
		return false, errors.New("el envio no esta en ruta")
	}

	//Agregamos la nueva parada al envio
	envioDB.Paradas = append(envioDB.Paradas, parada.GetParada().GetModel())

	//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
	return true, service.envioRepository.ActualizarEnvio(envioDB)
}

func (service *EnvioService) CambiarEstadoEnvio(envio *dto.Envio, usuario *dto.User) (bool, error) {
	//El estado deseado es el que se pasa con el objeto envio como parametro
	estadoDeseado := envio.Estado

	//Validamos el estado deseado
	if !model.EsUnEstadoEnvioValido(estadoDeseado) {
		return false, errors.New("el estado ingresado no es válido")
	}

	//Buscamos el envio en la base de datos para conocer el estado real
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envio.GetModel())

	if err != nil {
		return false, err
	}

	//Validamos el rol del usuario
	if !service.validarRol(usuario) {
		return false, errors.New("el usuario no tiene permisos para cambiar el estado del envio")
	}

	//Si el estado del envio no es compatible con el deseado, devolvemos un error
	if (estadoDeseado == model.EnRuta && envioDB.Estado != model.ADespachar) || (estadoDeseado == model.Despachado && envioDB.Estado != model.EnRuta) {
		return false, errors.New("el envio no puede pasar al estado " + fmt.Sprint(estadoDeseado) + " si esta en estado " + fmt.Sprint(envioDB.Estado))
	}

	//Actualizamos el envio en la base de datos
	envioDB.Estado = estadoDeseado
	err = service.envioRepository.ActualizarEnvio(envioDB)

	if err != nil {
		return false, err
	}

	//Si el envio pasa a estado Despachado, finaliza el viaje, por lo que hay que hacer otras operaciones
	if estadoDeseado == model.Despachado {
		service.finalizarViaje(dto.NewEnvio(*envioDB))
	}

	return true, nil
}

func (service *EnvioService) finalizarViaje(envio *dto.Envio) (bool, error) {
	//pasar pedidos a estado enviado
	err := service.entregarPedidosDeEnvio(envio)

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
	return service.pedidoRepository.ActualizarPedido(pedido)
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
	return service.productoRepository.ActualizarProducto(producto)
}

func (service *EnvioService) validarRol(usuario *dto.User) bool {
	return usuario.Rol == string(utils.Conductor)
}
