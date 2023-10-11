package handlers

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/services"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnvioHandler struct {
	envioService services.EnvioInterface
}

func NewEnvioHandler(envioService services.EnvioInterface) *EnvioHandler {
	return &EnvioHandler{envioService: envioService}
}

func (handler *EnvioHandler) ObtenerEnvios(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	envios, err := handler.envioService.ObtenerEnvios()

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

	envio, err := handler.envioService.ObtenerEnvioPorId(id)

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

	if operacion == false {
		log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
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

	if operacion == false {
		log.Printf("[handler:EnvioHandler][method:IniciarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
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

	if operacion == false {
		log.Printf("[handler:EnvioHandler][method:FinalizarViaje][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:FinalizarViaje][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}
