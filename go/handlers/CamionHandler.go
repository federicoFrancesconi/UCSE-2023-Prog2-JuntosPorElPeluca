package handlers

import (
	"TPIntegrador/dto"
	"TPIntegrador/services"
	"TPIntegrador/utils"
	"TPIntegrador/utils/logging"

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
		logging.LoggearErrorYResponder(c, "CamionHandler", "ObtenerCamiones", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "CamionHandler", "ObtenerCamiones", camiones, &user)
}

func (handler *CamionHandler) ObtenerCamionPorPatente(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	//Creamos el filtro
	filtro := utils.FiltroCamion{Patente: patente}

	camion, err := handler.camionService.ObtenerCamiones(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "ObtenerCamionPorPatente", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "CamionHandler", "ObtenerCamionPorPatente", camion, &user)
}

func (handler *CamionHandler) CrearCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	err := c.ShouldBindJSON(&camion)
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "CrearCamion", err, &user)
		return
	}

	//Si hay un error, lo devolvemos
	err = handler.camionService.CrearCamion(&camion, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "CrearCamion", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "CamionHandler", "CrearCamion", true, &user)
}

func (handler *CamionHandler) ActualizarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var camion dto.Camion
	err := c.ShouldBindJSON(&camion)
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "ActualizarCamion", err, &user)
	}

	//Pasamos el camion para actualizar al service
	err = handler.camionService.ActualizarCamion(&camion, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "ActualizarCamion", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "CamionHandler", "ActualizarCamion", true, &user)
}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	patente := c.Param("patente")

	//Generamos el objeto camion
	camionConPatente := dto.Camion{Patente: patente}

	//Si hay un error, lo devolvemos
	err := handler.camionService.EliminarCamion(&camionConPatente, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "CamionHandler", "EliminarCamion", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "CamionHandler", "EliminarCamion", true, &user)
}
