package handlers

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/services"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"log"
	"net/http"
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
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)

		// Devuelve badRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaCreacionFinStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01T00:00:00Z")
	fechaCreacionFin, err := time.Parse(time.RFC3339, fechaCreacionFinStr)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)

		// Devuelve badRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][cantidad:%d][user:%s]", len(pedidos), user.Codigo)

	c.JSON(http.StatusOK, pedidos)
}

func (handler *PedidoHandler) ObtenerCantidadPedidosPorEstado(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Obtenemos el array de cantidades del service
	cantidades, err := handler.pedidoService.ObtenerCantidadPedidosPorEstado()

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerCantidadesPedidosPorEstado][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:PedidoHandler][method:ObtenerCantidadesPedidosPorEstado][cantidad:%d][user:%s]", len(cantidades), user.Codigo)

	c.JSON(http.StatusOK, cantidades)
}

func (handler *PedidoHandler) CrearPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var pedido dto.Pedido

	//Parseamos el body del request y lo guardamos en el objeto pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		log.Printf("[handler:PedidoHandler][method:CrearPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creamos el pedido en la base de datos
	if err := handler.pedidoService.CrearPedido(&pedido, &user); err != nil {
		log.Printf("[handler:PedidoHandler][method:CrearPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:PedidoHandler][method:CrearPedido][id:%s][user:%s]", pedido.Id, user.Codigo)

	c.JSON(http.StatusOK, true)
}

func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//Generamos el objeto pedido
	pedido := dto.Pedido{Id: id}

	//Aceptamos el pedido
	if err := handler.pedidoService.AceptarPedido(&pedido); err != nil {
		log.Printf("[handler:PedidoHandler][method:AceptarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

func (handler *PedidoHandler) CancelarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//Generamos el objeto pedido
	pedido := dto.Pedido{Id: id}

	//Cancelamos el pedido
	if err := handler.pedidoService.CancelarPedido(&pedido); err != nil {
		log.Printf("[handler:PedidoHandler][method:CancelarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
