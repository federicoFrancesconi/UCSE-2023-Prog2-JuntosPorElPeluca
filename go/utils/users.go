package utils

import (
	"TPIntegrador/clients/responses"

	"github.com/gin-gonic/gin"
)

const (
    RolAdministrador = "Administrador"
    RolUsuario       = "Usuario"
    RolConductor     = "Conductor"
)

func SetUserInContext(c *gin.Context, user *responses.UserInfo) {
	c.Set("UserInfo", user)
}

func GetUserInfoFromContext(c *gin.Context) *responses.UserInfo {
	userInfo, _ := c.Get("UserInfo")

	user, _ := userInfo.(*responses.UserInfo)

	return user
}
