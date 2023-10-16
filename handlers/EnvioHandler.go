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

type EnvioHandler struct {
	envioService services.EnvioServiceInterface
}

func NewEnvioHandler(envioService services.EnvioServiceInterface) *EnvioHandler {
	return &EnvioHandler{envioService: envioService}
}

func (handler *EnvioHandler) ObtenerEnvios(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.DefaultQuery("patente", "")
	ultimaParada := c.DefaultQuery("ultimaParada", "")

	//Convierto el estado a integer para buscar el Estado en el "enum" de EstadoEnvio
	estadoStr := c.DefaultQuery("estado", "-1")
	estado, err := strconv.Atoi(estadoStr)

	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][envio:%+v][user:%s]", err.Error(), user.Codigo)

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

	//Llama al service
	envios, err := handler.envioService.ObtenerEnviosFiltrados(patente, model.EstadoEnvio(estado), ultimaParada, fechaCreacionComienzo, fechaCreacionFin)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:AulaHandler][method:ObtenerEnvios][cantidad:%d][user:%s]", len(envios), user.Codigo)

	c.JSON(http.StatusOK, envios)
}

func (handler *EnvioHandler) ObtenerEnvioPorId(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	//convertir id a int
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	envio, err := handler.envioService.ObtenerEnvioPorId(idInt)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][id:%s][user:%s]", id, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) CrearEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Si hay un error, lo devolvemos
	if err := handler.envioService.CrearEnvio(&envio); err != nil {
		log.Printf("[handler:EnvioHandler][method:CrearEnvio][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:CrearEnvio][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) AgregarParada(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operacion, err := handler.envioService.AgregarParada(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !operacion {
		log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //es correcto devolver bad request aca?
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) IniciarViaje(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operacion, err := handler.envioService.IniciarViaje(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:IniciarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !operacion {
		log.Printf("[handler:EnvioHandler][method:IniciarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:IniciarViaje][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) FinalizarViaje(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operacion, err := handler.envioService.FinalizarViaje(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:FinalizarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !operacion {
		log.Printf("[handler:EnvioHandler][method:FinalizarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:FinalizarViaje][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}
