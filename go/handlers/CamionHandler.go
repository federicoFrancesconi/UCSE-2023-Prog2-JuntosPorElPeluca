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
	camionService services.CamionServiceInterface
}

func NewCamionHandler(camionService services.CamionServiceInterface) *CamionHandler {
	return &CamionHandler{camionService: camionService}
}

func (handler *CamionHandler) ObtenerCamiones(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Creo un filtro vacio
	filtro := utils.FiltroCamion{}

	camiones, err := handler.camionService.ObtenerCamiones(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:CamionHandler][method:ObtenerCamiones][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:ObtenerCamiones][cantidad:%d][user:%s]", len(camiones), user.Codigo)

	c.JSON(http.StatusOK, camiones)
}

func (handler *CamionHandler) ObtenerCamionPorPatente(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	//Creamos el filtro
	filtro := utils.FiltroCamion{Patente: patente}

	camion, err := handler.camionService.ObtenerCamiones(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:CamionHandler][method:ObtenerCamionPorPatente][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:ObtenerCamionPorPatente][patente:%s][user:%s]", patente, user.Codigo)

	c.JSON(http.StatusOK, camion)
}

func (handler *CamionHandler) CrearCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Si hay un error, lo devolvemos
	if err := handler.camionService.CrearCamion(&camion, &user); err != nil {
		log.Printf("[handler:CamionHandler][method:CrearCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:CrearCamion][camion:%+v][user:%s]", camion, user.Codigo)

	c.JSON(http.StatusOK, true)
}

func (handler *CamionHandler) ActualizarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Si hay un error, lo devolvemos
	if err := handler.camionService.ActualizarCamion(&camion); err != nil {
		log.Printf("[handler:CamionHandler][method:ActualizarCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:ActualizarCamion][camion:%+v][user:%s]", camion, user.Codigo)

	c.JSON(http.StatusOK, true)
}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	//Generamos el objeto camion
	camionConPatente := dto.Camion{Patente: patente}

	//Si hay un error, lo devolvemos
	if err := handler.camionService.EliminarCamion(&camionConPatente); err != nil {
		log.Printf("[handler:CamionHandler][method:EliminarCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:CamionHandler][method:EliminarCamion][patente:%s][user:%s]", patente, user.Codigo)

	c.JSON(http.StatusOK, true)
}
