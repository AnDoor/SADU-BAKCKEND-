package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) LoginUserHandler(ctx *gin.Context) {
	var loginData = schema.LoginDTO{}
	err := ctx.BindJSON(loginData)
	token, err := u.service.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER", "Error logging in")
		return
	}
	helpers.SendSucces(ctx, "Successfully logged in", token)
}
