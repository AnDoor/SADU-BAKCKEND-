package config

import (
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

var DB *gorm.DB

func ConnectDB() {
	var error error

	DB, error = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	log.Default().Println("Token or Database URL is empty. Using local database.")

	if error != nil {
		log.Fatal(error)
	}
}

func SyncDB() error {
	models := []any{
		&schema.University{},
		&schema.Athlete{},
		&schema.Major{},
		&schema.Discipline{},
		&schema.Event{},
		&schema.Tourney{},
		&schema.Teacher{},
		&schema.Team{},
		&schema.AthleteDiscipline{},
		&schema.AthleteEvent{},
		&schema.AthleteTeam{},
		&schema.TeacherDiscipline{},
	}

	DB.AutoMigrate(models...)

	if err := DB.SetupJoinTable(&schema.Athlete{}, "Disciplines", &schema.AthleteDiscipline{}); err != nil {
		return err
	}
	log.Println("Setup AthleteDiscilpines seeded successfully")

	if err := DB.SetupJoinTable(&schema.Athlete{}, "Teams", &schema.AthleteTeam{}); err != nil {
		return err
	}
	log.Println("Setup AthleteTeam seeded successfully")

	if err := DB.SetupJoinTable(&schema.Teacher{}, "Disciplines", &schema.TeacherDiscipline{}); err != nil {
		return err
	}
	log.Println("Setup TeacherDisciplines seeded successfully")

	if err := DB.SetupJoinTable(&schema.Athlete{}, "Events", &schema.AthleteEvent{}); err != nil {
		return err
	}
	log.Println("Setup AthleteEvent seeded successfully")
	return nil
}
