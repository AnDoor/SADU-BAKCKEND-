package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
	"uneg.edu.ve/servicio-sadu-back/internal/routes"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	
	r := gin.Default()
	//configuracion de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://dominio.uneg.edu.ve"}, // "https://dominio.uneg.edu.ve" es cuando tengamos algun dominio ya puesto
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	athleteService := services.AthleteService{DB: config.DB}
	athleteHandler := handlers.NewAthleteHandler(&athleteService)

	universityService := services.UniversityServices{DB: config.DB}
	universityHandler := handlers.NewUniversityHandler(&universityService)

	disciplineService := services.DisciplineServices{DB: config.DB}
	disciplineHandler := handlers.NewDisciplineHandler(&disciplineService)

	majorService := services.MajorServices{DB: config.DB}
	majorHandler := handlers.NewMajorHandler(&majorService)

	tourneyService := services.TourneyServices{DB: config.DB}
	tourneyHandler := handlers.NewTourneyHandler(&tourneyService)

	teacherService := services.TeacherService{DB: config.DB}
	teacherHandler := handlers.NewTeacherHandler(&teacherService)

	teamService := services.TeamServices{DB: config.DB}
	teamHandler := handlers.NewTeamHandler(&teamService)

	eventService := services.EventService{DB: config.DB}
	eventHandlers := handlers.NewEventHandler(&eventService)

	/*rutas*/
	routes.RegisterAthletesRoutes(r.Group("/athletes"), athleteHandler)
	routes.RegisterUniversityRoutes(r.Group("/universities"), universityHandler)
	routes.RegisterDisciplines(r.Group("/disciplines"), disciplineHandler)
	routes.RegisterMajorsRoutes(r.Group("/majors"), majorHandler)
	routes.RegisterTourney(r.Group("/tourneys"), tourneyHandler)
	routes.RegisterTeacherRoutes(r.Group("/teachers"), teacherHandler)
	routes.RegisterTeamRoutes(r.Group("/teams"), teamHandler)
	routes.RegisterEventsRouters(r.Group("/events"), eventHandlers)
	log.Println(" Server corriendo en http://localhost:8080")
	r.Run(":8080")
	println("Exitted")
}
