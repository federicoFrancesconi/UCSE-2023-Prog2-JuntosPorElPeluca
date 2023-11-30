package logging

import (
	"TPIntegrador/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggearErrorYResponder(c *gin.Context, handler string, metodo string, err error, user *dto.User) {
	log.Printf("[handler:%s][método:%s][error:%s][user:%s]", handler, metodo, err.Error(), user.Codigo)

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func LoggearResultadoYResponder(c *gin.Context, handler string, metodo string, result interface{}, user *dto.User) {
	log.Printf("[handler:%s][método:%s][exitoso][user:%s]", handler, metodo, user.Codigo)

	c.JSON(http.StatusOK, result)
}
