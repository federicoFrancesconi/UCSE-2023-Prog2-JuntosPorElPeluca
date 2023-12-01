package handlers

import (
	"TPIntegrador/dto"
	"TPIntegrador/services"
	"TPIntegrador/utils"
	"TPIntegrador/utils/logging"
	"TPIntegrador/model"
	"time"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	pedidoService services.PedidoServiceInterface
}

func NewPedidoHandler(pedidoService services.PedidoServiceInterface) *PedidoHandler {
	return &PedidoHandler{pedidoService: pedidoService}
}

func (handler *PedidoHandler) ObtenerPedidos(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Obtenemos el id del envio, si es que se filtró por el mismo
	idEnvio := c.DefaultQuery("idEnvio", "")

	//Convierto el estado a integer para buscar el Estado en el "enum" de EstadoPedido
	estado := c.DefaultQuery("estado", "")

	// Convierte las fechas string a time.Time
	fechaCreacionComienzoStr := c.DefaultQuery("fechaCreacionComienzo", "0001-01-01T00:00:00Z")
	fechaCreacionComienzo, err := time.Parse(time.RFC3339, fechaCreacionComienzoStr)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "ObtenerPedidos", err, &user)
		return
	}

	fechaCreacionFinStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01T00:00:00Z")
	fechaCreacionFin, err := time.Parse(time.RFC3339, fechaCreacionFinStr)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "ObtenerPedidos", err, &user)
		return
	}

	//Creamos el filtro con los datos obtenidos
	filtro := utils.FiltroPedido{
		IdEnvio:               idEnvio,
		Estado:                model.EstadoPedido(estado),
		FechaCreacionComienzo: fechaCreacionComienzo,
		FechaCreacionFin:      fechaCreacionFin,
	}

	//Obtenemos los pedidos
	pedidos, err := handler.pedidoService.ObtenerPedidosFiltrados(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "ObtenerPedidos", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "PedidoHandler", "ObtenerPedidos", pedidos, &user)
}

func (handler *PedidoHandler) ObtenerCantidadPedidosPorEstado(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Obtenemos el array de cantidades del service
	cantidades, err := handler.pedidoService.ObtenerCantidadPedidosPorEstado()

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "ObtenerCantidadPedidosPorEstado", err, &user)
		return
	}

	logging.LoggearResultadoYResponder(c, "PedidoHandler", "ObtenerCantidadPedidosPorEstado", cantidades, &user)
}

func (handler *PedidoHandler) CrearPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var pedido dto.Pedido

	//Parseamos el body del request y lo guardamos en el objeto pedido
	err := c.ShouldBindJSON(&pedido)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "CrearPedido", err, &user)
		return
	}

	//Creamos el pedido en la base de datos
	err = handler.pedidoService.CrearPedido(&pedido, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "CrearPedido", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "PedidoHandler", "CrearPedido", true, &user)
}

func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//Generamos el objeto pedido
	pedido := dto.Pedido{Id: id}

	//Aceptamos el pedido
	err := handler.pedidoService.AceptarPedido(&pedido, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "AceptarPedido", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "PedidoHandler", "AceptarPedido", true, &user)
}

func (handler *PedidoHandler) CancelarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//Generamos el objeto pedido
	pedido := dto.Pedido{Id: id}

	//Cancelamos el pedido
	err := handler.pedidoService.CancelarPedido(&pedido, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "CancelarPedido", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "PedidoHandler", "CancelarPedido", true, &user)
}