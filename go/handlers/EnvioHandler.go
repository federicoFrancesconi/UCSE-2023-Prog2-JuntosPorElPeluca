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
	fechaCreacionDesdeStr := c.DefaultQuery("fechaCreacionComienzo", "0001-01-01T00:00:00Z")
	fechaCreacionDesde, err := time.Parse(time.RFC3339, fechaCreacionDesdeStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaCreacionDesde = time.Time{}
	}

	fechaCreacionHastaStr := c.DefaultQuery("fechaCreacionFin", "0001-01-01T00:00:00Z")
	fechaCreacionHasta, err := time.Parse(time.RFC3339, fechaCreacionHastaStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaCreacionHasta = time.Time{}
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

	envio, err := handler.envioService.ObtenerEnvioPorId(&dto.Envio{Id: id})

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
	fechaDesdeStr := c.DefaultQuery("fechaDesde", "0001-01-01T00:00:00Z")
	fechaDesde, err := time.Parse(time.RFC3339, fechaDesdeStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaDesde = time.Time{}
	}

	fechaHastaStr := c.DefaultQuery("fechaHasta", "0001-01-01T00:00:00Z")
	fechaHasta, err := time.Parse(time.RFC3339, fechaHastaStr)
	if err != nil {
		// Si hay un error en el parseo, devuelve una fecha default
		fechaHasta = time.Time{}
	}

	//TODO: probar si anda el filtro solamente pasandole un par de parametros
	//Creamos el filtro, que tiene en cuenta solamente las fechas
	filtro := utils.FiltroEnvio{
		FechaUltimaActualizacionDesde: fechaDesde,
		FechaUltimaActualizacionHasta: fechaHasta,
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

	c.JSON(http.StatusOK, true)
}

func (handler *EnvioHandler) AgregarParada(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el id como parametro
	id := c.Param("id")

	//Obtenemos la nueva parada
	var parada dto.Parada
	if err := c.ShouldBindJSON(&parada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creamos el envio para pasarle al service
	envio := dto.Envio{
		Id: id,
		Paradas: []dto.Parada{
			parada,
		},
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

	c.JSON(http.StatusOK, true)
}

func (handler *EnvioHandler) CambiarEstadoEnvio(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el id del envio a finalizar
	id := c.Param("id")

	//Recibimos el estado deseado para el envio
	estadoStr := c.DefaultQuery("estado", "-1")

	//Convierto el estado a integer para buscar el Estado en el "enum" de EstadoEnvio
	estado, err := strconv.Atoi(estadoStr)

	if err != nil {
		log.Printf("[handler:EnvioHandler][method:CambiarEstadoEnvio][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creamos el envio para pasarle al service
	envio := dto.Envio{
		Id:     id,
		Estado: model.EstadoEnvio(estado),
	}

	operacion, err := handler.envioService.CambiarEstadoEnvio(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:CambiarEstadoEnvio][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !operacion {
		log.Printf("[handler:EnvioHandler][method:CambiarEstadoEnvio][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:EnvioHandler][method:CambiarEstadoEnvio][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, true)
}
