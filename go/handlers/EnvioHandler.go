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

	//Creamos el filtro
	filtro := utils.FiltroEnvio{
		PatenteCamion:         patente,
		Estado:                model.EstadoEnvio(estado),
		UltimaParada:          ultimaParada,
		FechaCreacionComienzo: fechaCreacionComienzo,
		FechaCreacionFin:      fechaCreacionFin,
	}

	//Llama al service
	envios, err := handler.envioService.ObtenerEnviosFiltrados(filtro)

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

	envio, err := handler.envioService.ObtenerEnvioPorId(&dto.Envio{Id: idInt})

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

func (handler *EnvioHandler) ObtenerBeneficioEntreFechas(c *gin.Context) {
	// Convierte las fechas string a time.Time
	fechaDesdeStr := c.DefaultQuery("fechaCreacionComienzo", "0001-01-01T00:00:00Z")
	fechaDesde, err := time.Parse(time.RFC3339, fechaDesdeStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaDesde = time.Time{}
	}

	fechaHastaStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01T00:00:00Z")
	fechaHasta, err := time.Parse(time.RFC3339, fechaHastaStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaHasta = time.Time{}
	}

	//Creamos el filtro
	filtro := utils.FiltroEnvio{
		PatenteCamion:         "",
		Estado:                -1,
		UltimaParada:          "",
		FechaCreacionComienzo: fechaDesde,
		FechaCreacionFin:      fechaHasta,
	}

	//Llama al service
	beneficio, err := handler.envioService.ObtenerBeneficioEntreFechas(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerBeneficioEntreFechas][envio:%+v]", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:ObtenerBeneficioEntreFechas][beneficio:%f]", beneficio)

	c.JSON(http.StatusOK, beneficio)
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
