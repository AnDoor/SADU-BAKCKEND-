package helpers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProblemDetails struct {
	Type     string `json:"type"`
	Tittle   string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func SendError(ctx *gin.Context, code int, title, detail string) {
	problem := ProblemDetails{
		Type:     "about:blank",
		Tittle:   title,
		Status:   code,
		Detail:   detail,
		Instance: ctx.Request.URL.Path,
	}
	ctx.JSON(code, problem)
}

func SendSucces(ctx *gin.Context, operation string, data any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfully", operation),
		"data":    data,
	})
}
