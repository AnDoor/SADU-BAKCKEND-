package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
	"uneg.edu.ve/servicio-sadu-back/internal/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	handlers.DB = config.DB

	r := gin.Default()
	r.Use(cors.Default())

	routes.RegisterAthletesRoutes(r)
	routes.RegisterTeamsRoutes(r)
	routes.RegisterTeachersRoutes(r)
	routes.RegisterEventsRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/athletes/:id", func (c *gin.Context){
		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})

	
	r.Run(":8080")
	println("Exitted")
}
