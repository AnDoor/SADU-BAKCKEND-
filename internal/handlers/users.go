package handlers

import (
<<<<<<< HEAD
	"fmt"
=======
	"log"
>>>>>>> 8eb16fdf2d4640cf06d3802b5c9262f491fdf7f7
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
	var loginData schema.LoginDTO
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		helpers.SendError(ctx, http.StatusBadRequest, "INVALID_INPUT", "Invalid login data")
		return
	}

	token, err := u.service.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		helpers.SendError(ctx, http.StatusUnauthorized, "AUTH_FAILED", "Invalid credentials")
		return
	}
	fmt.Println("------------------------------------------")
	fmt.Println("🔑 TOKEN PARA POSTMAN:")
	fmt.Println(token)
	fmt.Println("------------------------------------------")
	helpers.SendSucces(ctx, "Successfully logged in", token)
}
