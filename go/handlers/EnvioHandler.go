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

	estado := c.DefaultQuery("estado", "")

	// Convierte las fechas string a time.Time
	fechaCreacionDesdeStr := c.DefaultQuery("fechaCreacionComienzo", "0001-01-01")
	fechaCreacionDesde, err := time.Parse("2006-01-02", fechaCreacionDesdeStr)
	//Contemplamos si hay errores en el parseo
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerEnvios", err, &user)
	}

	fechaCreacionHastaStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01")
	fechaCreacionHasta, err := time.Parse("2006-01-02", fechaCreacionHastaStr)
	//Contemplamos si hay errores en el parseo
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerEnvios", err, &user)
	}

	//Creamos el filtro
	filtro := utils.FiltroEnvio{
		PatenteCamion:      patente,
		Estado:             model.EstadoEnvio(estado),
		UltimaParada:       ultimaParada,
		FechaCreacionDesde: fechaCreacionDesde,
		FechaCreacionHasta: fechaCreacionHasta,
	}

	//Llama al service
	envios, err := handler.envioService.ObtenerEnvios(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerEnvios", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "ObtenerEnvios", envios, &user)
}

func (handler *EnvioHandler) ObtenerEnvioPorId(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	id := c.Param("id")

	envio, err := handler.envioService.ObtenerEnvioPorId(&dto.Envio{Id: id})

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerEnvioPorId", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "ObtenerEnvioPorId", envio, &user)
}

func (handler *EnvioHandler) ObtenerBeneficioEntreFechas(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Convierte las fechas string a time.Time
	fechaDesdeStr := c.DefaultQuery("fechaDesde", "0001-01-01")
	fechaDesde, err := time.Parse("2006-01-02", fechaDesdeStr)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerBeneficioEntreFechas", err, &user)
		return
	}

	//Convertimos la fecha a un time con hora 0
	fechaDesde = time.Date(fechaDesde.Year(), fechaDesde.Month(), fechaDesde.Day(), 0, 0, 0, 0, fechaDesde.Location())

	fechaHastaStr := c.DefaultQuery("fechaHasta", "0001-01-01")
	fechaHasta, err := time.Parse("2006-01-02", fechaHastaStr)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerBeneficioEntreFechas", err, &user)
		return
	}

	//Convertimos la fecha a un time con hora 0
	fechaHasta = time.Date(fechaHasta.Year(), fechaHasta.Month(), fechaHasta.Day(), 0, 0, 0, 0, fechaHasta.Location())

	//Creamos el filtro, que tiene en cuenta solamente las fechas
	filtro := utils.FiltroEnvio{
		FechaUltimaActualizacionDesde: fechaDesde,
		FechaUltimaActualizacionHasta: fechaHasta,
	}

	//Llama al service
	beneficioTemporal, err := handler.envioService.ObtenerBeneficioTemporal(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerBeneficioEntreFechas", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "ObtenerBeneficioEntreFechas", beneficioTemporal, &user)
}

func (handler *EnvioHandler) ObtenerCantidadEnviosPorEstado(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Obtenemos el array de cantidades del service
	cantidades, err := handler.envioService.ObtenerCantidadEnviosPorEstado()

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerCantidadEnviosPorEstado", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "ObtenerCantidadEnviosPorEstado", cantidades, &user)
}

func (handler *EnvioHandler) CrearEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var envio dto.Envio
	err := c.ShouldBindJSON(&envio)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "CrearEnvio", err, &user)
		return
	}

	//Si hay un error, lo devolvemos
	err = handler.envioService.CrearEnvio(&envio, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "CrearEnvio", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "CrearEnvio", true, &user)
}

func (handler *EnvioHandler) AgregarParada(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Obtenemos la nueva parada
	var parada dto.NuevaParada
	err := c.ShouldBindJSON(&parada)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "AgregarParada", err, &user)
		return
	}

	operacion, err := handler.envioService.AgregarParada(&parada, &user)
	if err != nil || !operacion {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "AgregarParada", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "AgregarParada", true, &user)
}

func (handler *EnvioHandler) CambiarEstadoEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el envio en el body
	//Este contiene el id del envio y el nuevo estado
	var envio dto.Envio
	err := c.ShouldBindJSON(&envio)
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "CambiarEstadoEnvio", err, &user)
		return
	}

	operacion, err := handler.envioService.CambiarEstadoEnvio(&envio, &user)
	if err != nil || !operacion {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "CambiarEstadoEnvio", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "CambiarEstadoEnvio", true, &user)
}
