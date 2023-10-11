package handlers

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/services"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CamionHandler struct {
	camionService services.CamionInterface
}

func NewCamionHandler(camionService services.CamionInterface) *CamionHandler {
	return &CamionHandler{camionService: camionService}
}

func (handler *CamionHandler) ObtenerCamiones(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	camiones := handler.camionService.ObtenerCamiones()

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:AulaHandler][method:ObtenerAulas][cantidad:%d][user:%s]", len(camiones), user.Codigo)

	c.JSON(http.StatusOK, camiones)
}

func (handler *CamionHandler) ObtenerCamionPorPatente(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	camion := handler.camionService.ObtenerCamionPorPatente(patente)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:ObtenerCamionPorId][patente:%s][user:%s]", patente, user.Codigo)

	c.JSON(http.StatusOK, camion)
}

func (handler *CamionHandler) CrearCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.camionService.CrearCamion(&camion)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:CrearCamion][camion:%+v][user:%s]", camion, user.Codigo)

	c.JSON(http.StatusOK, camion)
}

func (handler *CamionHandler) ActualizarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.camionService.ActualizarCamion(&camion)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:ActualizarCamion][camion:%+v][user:%s]", camion, user.Codigo)

	c.JSON(http.StatusOK, camion)
}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	handler.camionService.EliminarCamion(patente)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:EliminarCamion][patente:%s][user:%s]", patente, user.Codigo)

	c.JSON(http.StatusOK, true)
}
