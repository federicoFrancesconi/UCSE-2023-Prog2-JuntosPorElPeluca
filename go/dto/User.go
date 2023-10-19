package dto

import "UCSE-2023-Prog2-TPIntegrador/clients/responses"

type User struct {
	Codigo   string `json:"codigo"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
}

func NewUser(userInfo *responses.UserInfo) User {
	user := User{}
	if userInfo != nil {
		user.Codigo = userInfo.Codigo
		user.Email = userInfo.Email
		user.Username = userInfo.Username
		user.Rol = userInfo.Rol
	}
	return user
}