package handlers

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/services"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"log"
	"net/http"
	"strconv"
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

	idEnvioStr := c.DefaultQuery("idEnvio", "0")
	idEnvio, err := strconv.Atoi(idEnvioStr)

	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Convierto el estado a integer para buscar el Estado en el "enum" de EstadoPedido
	estadoStr := c.DefaultQuery("estado", "-1")
	estado, err := strconv.Atoi(estadoStr)

	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convierte las fechas string a time.Time
	fechaCreacionComienzoStr := c.DefaultQuery("fechaCreacionComienzo", "0001-01-01T00:00:00Z")
	fechaCreacionComienzo, err := time.Parse(time.RFC3339, fechaCreacionComienzoStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaCreacionComienzo = time.Time{}
	}

	fechaCreacionFinStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01T00:00:00Z")
	fechaCreacionFin, err := time.Parse(time.RFC3339, fechaCreacionFinStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaCreacionFin = time.Time{}
	}

	pedidos, err := handler.pedidoService.ObtenerPedidosFiltrados(idEnvio, model.EstadoPedido(estado), fechaCreacionComienzo, fechaCreacionFin)

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
	if err := handler.pedidoService.CrearPedido(&pedido); err != nil {
		log.Printf("[handler:PedidoHandler][method:CrearPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pedido)
}

func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//convertir id a int
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[handler:PedidoHandler][method:EnviarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Aceptamos el pedido
	if err := handler.pedidoService.AceptarPedido(idInt); err != nil {
		log.Printf("[handler:PedidoHandler][method:AceptarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO: capaz este mal que devuelva nil
	c.JSON(http.StatusOK, nil)
}

func (handler *PedidoHandler) CancelarPedido(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//convertir id a int
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[handler:PedidoHandler][method:EnviarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Cancelamos el pedido
	if err := handler.pedidoService.CancelarPedido(idInt); err != nil {
		log.Printf("[handler:PedidoHandler][method:CancelarPedido][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: capaz este mal que devuelva nil
	c.JSON(http.StatusOK, nil)
}
