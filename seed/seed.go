package main

import (
	"log"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
	"uneg.edu.ve/servicio-sadu-back/internal/routes"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
	. "uneg.edu.ve/servicio-sadu-back/schema"

	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	db := config.DB

	if err := seedDatabase(db); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	} else {
		log.Println("Database seeded successfully")
	}

	athleteService := services.AthleteService{DB: db}
	athleteHandler := handlers.NewAthleteHandler(&athleteService)

	universityService := services.UniversityServices{DB: db}
	universityHandler := handlers.NewUniversityHandler(&universityService)

	disciplineService := services.DisciplineServices{DB: db}
	disciplineHandler := handlers.NewDisciplineHandler(&disciplineService)

	majorService := services.MajorServices{DB: db}
	majorHandler := handlers.NewMajorHandler(&majorService)

	tourneyService := services.TourneyServices{DB: db}
	tourneyHandler := handlers.NewTourneyHandler(&tourneyService)

	teacherService := services.TeacherService{DB: db}
	teacherHandler := handlers.NewTeacherHandler(&teacherService)

	r := gin.Default()
	//configuracion de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "https://dominio.uneg.edu.ve"}, // "https://dominio.uneg.edu.ve" es cuando tengamos algun dominio ya puesto
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	/*rutas*/
	routes.RegisterAthletesRoutes(r.Group("/athletes"), athleteHandler)
	routes.RegisterUniversityRoutes(r.Group("/university"), universityHandler)
	routes.RegisterDisciplines(r.Group("/discipline"), disciplineHandler)
	routes.RegisterMajorsRoutes(r.Group("/major"), majorHandler)
	routes.RegisterTourney(r.Group("/tourney"), tourneyHandler)
	routes.RegisterTeacherRoutes(r.Group("/teachers"), teacherHandler)

	log.Println(" Server corriendo en http://localhost:8080")
	r.Run(":8080")
}

func seedDatabase(db *gorm.DB) error {
	log.Println("Seeding database...")
	faker := gofakeit.New(777)

	// Seed Majors
	carreras := []Major{
		{Name: "Informatica"},
		{Name: "Educación"},
		{Name: "Industrial"},
		{Name: "Medicina"},
		{Name: "Derecho"},
		{Name: "Ingeniería Civil"},
	}
	if err := db.Create(&carreras).Error; err != nil {
		return err
	}
	log.Println("Majors seeded successfully")

	// Seed Universities
	universidades := []University{
		{Name: "UNEG Villa Asia", Local: true},
		{Name: "UNEG Maturin", Local: true},
		{Name: "UCAB", Local: false},
		{Name: "UCV", Local: false},
		{Name: "USB", Local: false},
	}
	if err := db.Create(&universidades).Error; err != nil {
		return err
	}
	log.Println("Universities seeded successfully")

	// Seed Disciplines
	disciplinas := []Discipline{
		{Name: "Volleybol"},
		{Name: "Futbol"},
		{Name: "Basquetbol"},
		{Name: "Natación"},
		{Name: "Atletismo"},
		{Name: "Tenis"},
	}
	if err := db.Create(&disciplinas).Error; err != nil {
		return err
	}
	log.Println("Disciplines seeded successfully")

	// Seed Teams
	equipos := []Team{}
	genders := []Gender{GenderM, GenderF}
	for i, universidad := range universidades {
		for j, disciplina := range disciplinas {
			for _, gender := range genders {
				equipos = append(equipos, Team{
					Name:         faker.PetName(),
					Regular:      faker.Bool(),
					Category:     gender,
					DisciplineID: schema.RegularIDs(disciplina.ID),
					UniversityID: schema.RegularIDs(universidad.ID),
				})
				// Add some irregular teams
				if i%2 == 0 && j%2 == 0 {
					equipos = append(equipos, Team{
						Name:         faker.PetName() + " Reserva",
						Regular:      false,
						Category:     gender,
						DisciplineID: schema.RegularIDs(disciplina.ID),
						UniversityID: schema.RegularIDs(universidad.ID),
					})
				}
			}
		}
	}
	if err := db.Create(&equipos).Error; err != nil {
		return err
	}
	log.Println("Teams seeded successfully")

	// Seed Athletes
	atletas := []Athlete{}
	for i := range 50 {
		GovId := strconv.Itoa(faker.IntRange(28398421, 31613490) + i)
		GovId = "V-" + GovId
		gender := genders[i%2]
		majorIndex := i % len(carreras)

		atletas = append(atletas, Athlete{
			FirstNames:      faker.FirstName(),
			LastNames:       faker.LastName() + " " + faker.LastName(),
			PhoneNum:        faker.Phone(),
			Inscripted:      faker.Bool(),
			Email:           faker.Email(),
			Gender:          gender,
			InscriptionDate: faker.DateRange(time.Now().AddDate(-2, 0, 0), time.Now()),
			Regular:         faker.Bool(),
			GovID:           GovId,
			MajorID:         schema.RegularIDs(carreras[majorIndex].ID),
		})
	}
	if err := db.Create(&atletas).Error; err != nil {
		return err
	}
	log.Println("Athletes seeded successfully")

	// Seed Teachers
	profesores := []Teacher{}
	for i := range 15 {
		GovId := strconv.Itoa(faker.IntRange(15000000, 25000000) + i)
		GovId = "V-" + GovId
		profesores = append(profesores, Teacher{
			FirstNames: faker.FirstName(),
			LastNames:  faker.LastName() + " " + faker.LastName(),
			PhoneNum:   faker.Phone(),
			Email:      faker.Email(),
			GovID:      GovId,
		})
	}
	if err := db.Create(&profesores).Error; err != nil {
		return err
	}
	log.Println("Teachers seeded successfully")

	// Seed Tourneys
	torneos := []Tourney{
		{Name: "Juegos Interuniversitarios Nacional"},
		{Name: "Copa Universitaria Regional"},
		{Name: "Torneo de Verano"},
		{Name: "Campeonato Nacional Universitario"},
		{Name: "Liga Universitaria"},
	}
	if err := db.Create(&torneos).Error; err != nil {
		return err
	}
	log.Println("Tourneys seeded successfully")

	// Seed Events
	eventos := []Event{}
	statuses := []string{"Programado", "En Curso", "Finalizado", "Cancelado", "Pospuesto"}
	locations := []string{"Estadio Principal", "Cancha Norte", "Piscina Universitaria", "Gimnasio Central", "Campo de Atletismo"}

	for i := range 30 {
		homeTeamIndex := i % len(equipos)
		oppositeTeamIndex := (i + 1) % len(equipos)
		// Ensure teams are different
		if homeTeamIndex == oppositeTeamIndex {
			oppositeTeamIndex = (oppositeTeamIndex + 1) % len(equipos)
		}

		tourneyIndex := i % len(torneos)
		teacherIndex := i % len(profesores)
		disciplineIndex := i % len(disciplinas)

		eventos = append(eventos, Event{
			Name:                 faker.Sentence(3),
			Date:                 faker.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(1, 0, 0)),
			Status:               statuses[i%len(statuses)],
			Observation:          faker.Sentence(8),
			Ubication:            locations[i%len(locations)],
			HomePoints:           uint8(faker.IntRange(0, 100)),
			OppositePoints:       uint8(faker.IntRange(0, 100)),
			HomeTeamID:           schema.RegularIDs(equipos[homeTeamIndex].ID),
			OppositeTeamID:       schema.RegularIDs(equipos[oppositeTeamIndex].ID),
			TourneyID:            schema.RegularIDs(torneos[tourneyIndex].ID),
			ResponsableTeacherID: schema.RegularIDs(profesores[teacherIndex].ID),
			DisciplineID:         schema.RegularIDs(disciplinas[disciplineIndex].ID),
		})
	}
	if err := db.Create(&eventos).Error; err != nil {
		return err
	}
	log.Println("Events seeded successfully")

	// Setup and seed join tables
	if err := db.SetupJoinTable(&Athlete{}, "Disciplines", &AthleteDiscipline{}); err != nil {
		return err
	}
	log.Println("Setup AthleteDisciplines successfully")

	if err := db.SetupJoinTable(&Athlete{}, "Teams", &AthleteTeam{}); err != nil {
		return err
	}
	log.Println("Setup AthleteTeam successfully")

	if err := db.SetupJoinTable(&Teacher{}, "Disciplines", &TeacherDiscipline{}); err != nil {
		return err
	}
	log.Println("Setup TeacherDisciplines successfully")

	if err := db.SetupJoinTable(&Athlete{}, "Events", &AthleteEvent{}); err != nil {
		return err
	}
	log.Println("Setup AthleteEvent successfully")

	// Seed AthleteDiscipline relationships
	athleteDisciplines := []AthleteDiscipline{}
	for _, atleta := range atletas {
		// Each athlete practices 1-3 disciplines
		numDisciplines := faker.IntRange(1, 4)
		usedDisciplines := make(map[uint]bool)

		for range numDisciplines {
			disciplineIndex := faker.IntRange(0, len(disciplinas)-1)
			// Avoid duplicates
			if usedDisciplines[disciplinas[disciplineIndex].ID] {
				continue
			}
			usedDisciplines[disciplinas[disciplineIndex].ID] = true

			athleteDisciplines = append(athleteDisciplines, AthleteDiscipline{
				AthleteID:    schema.RegularIDs(atleta.ID),
				DisciplineID: schema.RegularIDs(disciplinas[disciplineIndex].ID),
				Regular:      faker.Bool(),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})
		}
	}
	if err := db.Create(&athleteDisciplines).Error; err != nil {
		return err
	}
	log.Println("AthleteDiscipline relationships seeded successfully")

	// Seed AthleteTeam relationships
	athleteTeams := []AthleteTeam{}
	for _, atleta := range atletas {
		// Each athlete is in 1-2 teams
		numTeams := faker.IntRange(1, 3)
		usedTeams := make(map[uint]bool)

		for range numTeams {
			teamIndex := faker.IntRange(0, len(equipos)-1)
			// Avoid duplicates
			if usedTeams[equipos[teamIndex].ID] {
				continue
			}
			usedTeams[equipos[teamIndex].ID] = true

			startDate := faker.DateRange(time.Now().AddDate(-2, 0, 0), time.Now().AddDate(-1, 0, 0))
			endDate := faker.DateRange(startDate, time.Now().AddDate(1, 0, 0))

			athleteTeams = append(athleteTeams, AthleteTeam{
				AthleteID: schema.RegularIDs(atleta.ID),
				TeamID:    schema.RegularIDs(equipos[teamIndex].ID),
				StartDate: startDate,
				EndDate:   endDate,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}
	if err := db.Create(&athleteTeams).Error; err != nil {
		return err
	}
	log.Println("AthleteTeam relationships seeded successfully")

	// Seed TeacherDiscipline relationships
	teacherDisciplines := []TeacherDiscipline{}
	for _, profesor := range profesores {
		// Each teacher teaches 1-2 disciplines
		numDisciplines := faker.IntRange(1, 3)
		usedDisciplines := make(map[uint]bool)

		for range numDisciplines {
			disciplineIndex := faker.IntRange(0, len(disciplinas)-1)
			// Avoid duplicates
			if usedDisciplines[disciplinas[disciplineIndex].ID] {
				continue
			}
			usedDisciplines[disciplinas[disciplineIndex].ID] = true

			startDate := faker.DateRange(time.Now().AddDate(-3, 0, 0), time.Now().AddDate(-1, 0, 0))
			endDate := faker.DateRange(startDate, time.Now().AddDate(2, 0, 0))

			teacherDisciplines = append(teacherDisciplines, TeacherDiscipline{
				TeacherID:    schema.RegularIDs(profesor.ID),
				DisciplineID: schema.RegularIDs(disciplinas[disciplineIndex].ID),
				StartDate:    startDate,
				EndDate:      endDate,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})
		}
	}
	if err := db.Create(&teacherDisciplines).Error; err != nil {
		return err
	}
	log.Println("TeacherDiscipline relationships seeded successfully")

	// Seed AthleteEvent relationships
	athleteEvents := []AthleteEvent{}
	for _, evento := range eventos {
		// Each event has 5-15 participating athletes
		numAthletes := faker.IntRange(5, 16)
		usedAthletes := make(map[uint]bool)

		for j := range numAthletes {
			athleteIndex := faker.IntRange(0, len(atletas)-1)
			// Avoid duplicates
			if usedAthletes[atletas[athleteIndex].ID] {
				continue
			}
			usedAthletes[atletas[athleteIndex].ID] = true

			// Assign athlete to either home or opposite team
			var teamID schema.RegularIDs
			if j%2 == 0 {
				teamID = evento.HomeTeamID
			} else {
				teamID = evento.OppositeTeamID
			}

			athleteEvents = append(athleteEvents, AthleteEvent{
				AthleteID: schema.RegularIDs(atletas[athleteIndex].ID),
				EventID:   schema.RegularIDs(evento.ID),
				TeamID:    teamID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}
	if err := db.Create(&athleteEvents).Error; err != nil {
		return err
	}
	log.Println("AthleteEvent relationships seeded successfully")

	log.Println("All database seeding completed successfully!")
	return nil
}
