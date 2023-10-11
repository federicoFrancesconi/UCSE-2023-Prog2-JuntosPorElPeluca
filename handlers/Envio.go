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

	aulas := handler.envioService.ObtenerEnvios()

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:AulaHandler][method:ObtenerAulas][cantidad:%d][user:%s]", len(aulas), user.Codigo)

	c.JSON(http.StatusOK, aulas)
}

func (handler *EnvioHandler) ObtenerEnvioPorId(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	envio := handler.envioService.ObtenerEnvioPorId(id)

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

	handler.envioService.CrearEnvio(&envio)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:CrearEnvio][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) ActualizarEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.envioService.ActualizarEnvio(&envio)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:ActualizarEnvio][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) EliminarEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	handler.envioService.EliminarEnvio(id)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:EliminarEnvio][id:%s][user:%s]", id, user.Codigo)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
