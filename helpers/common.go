package helpers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func SendSucces(ctx *gin.Context, operation string, data any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", operation),
		"data":    data,
	})
}
