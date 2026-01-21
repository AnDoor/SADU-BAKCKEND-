package schema

import (
	"time"

	"gorm.io/gorm"
)

// Status represents the possible states of an item
type Gender string

const (
	GenderM Gender = "Masculino"
	GenderF Gender = "Femenino"
)

// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
type Major struct { // Carrera
	gorm.Model
	Name     string
	Athletes []Athlete `gorm:"foreignKey:MajorID"`
}

type University struct {
	gorm.Model
	Name  string
	Local bool
	Teams []Team `gorm:"foreignKey:UniversityID"`
}

type Team struct {
	gorm.Model
	Name          string
	Regular       bool
	Category      Gender `gorm:"index;type:gender"`
	DisciplineID  RegularIDs
	UniversityID  RegularIDs
	Discipline    Discipline     `gorm:"foreignKey:DisciplineID"`
	University    University     `gorm:"foreignKey:UniversityID"`
	AthleteEvents []AthleteEvent `gorm:"foreignKey:TeamID"`
	Athletes      []Athlete      `gorm:"many2many:athlete_teams;"`
}

type Athlete struct {
	gorm.Model
	FirstNames      string
	LastNames       string
	PhoneNum        string
	Email           string
	Inscripted 		bool 
	Gender          Gender    `gorm:"index;type:gender"`
	InscriptionDate time.Time // Fecha de inscripcion
	Regular         bool      // Titular
	GovID           string    // Cedula
	MajorID         RegularIDs
	Teams           []Team       `gorm:"many2many:athlete_teams;"`
	Events          []Event      `gorm:"many2many:athlete_events;"` // foreignKey:AthleteID
	Disciplines     []Discipline `gorm:"many2many:athlete_disciplines;"`
}

//disciplina deportiva 
type Discipline struct {
	gorm.Model
	Name     string
	Teams    []Team    `gorm:"foreignKey:DisciplineID"`
	Events   []Event   `gorm:"foreignKey:DisciplineID"`
	Athletes []Athlete `gorm:"many2many:athlete_disciplines;"`
	Teachers []Teacher `gorm:"many2many:teacher_disciplines;"`
}

type Teacher struct {
	gorm.Model
	FirstNames  string
	LastNames   string
	PhoneNum    string
	Email       string
	GovID       string
	Events      []Event      `gorm:"foreignKey:ResponsableTeacherID"`
	Disciplines []Discipline `gorm:"many2many:teacher_disciplines;"`
}

type Tourney struct {
	gorm.Model
	Name   string
	Events []Event `gorm:"foreignKey:TourneyID"`
}

type Event struct {
	gorm.Model
	Name                 string
	Date                 time.Time
	Status               string
	Observation          string
	Ubication            string
	HomePoints           uint8
	OppositePoints       uint8
	HomeTeamID           RegularIDs
	OppositeTeamID       RegularIDs
	TourneyID            RegularIDs
	ResponsableTeacherID RegularIDs
	DisciplineID         RegularIDs
	HomeTeam             Team       `gorm:"foreignKey:HomeTeamID"`
	OppositeTeam         Team       `gorm:"foreignKey:OppositeTeamID"`
	Tourney              Tourney    `gorm:"foreignKey:TourneyID"`
	ResponsableTeacher   Teacher    `gorm:"foreignKey:ResponsableTeacherID"`
	Discipline           Discipline `gorm:"foreignKey:DisciplineID"`
	Athletes             []Athlete  `gorm:"many2many:athlete_events;"`
}

type AthleteDiscipline struct {
	AthleteID    RegularIDs `gorm:"primaryKey;"`
	DisciplineID RegularIDs `gorm:"primaryKey;"`
	Regular      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"index"`
	DeletedAt    gorm.DeletedAt
}

// err := db.SetupJoinTable(&Athlete{}, "Discipline", &AthleteDisciplines{})

type AthleteTeam struct {
	AthleteID RegularIDs `gorm:"primaryKey"`
	TeamID    RegularIDs `gorm:"primaryKey"`
	StartDate time.Time
	EndDate   time.Time `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// err := db.SetupJoinTable(&Athlete{}, "Team", &AthleteTeam{})

type TeacherDiscipline struct {
	TeacherID    RegularIDs `gorm:"primaryKey"`
	DisciplineID RegularIDs `gorm:"primaryKey"`
	StartDate    time.Time
	EndDate      time.Time `gorm:"index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

// err := db.SetupJoinTable(&Teacher{}, "Discipline", &TeacherDisciplines{})

type AthleteEvent struct {
	AthleteID RegularIDs `gorm:"primaryKey"`
	EventID   RegularIDs `gorm:"primaryKey;index"`
	TeamID    RegularIDs
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// err := db.SetupJoinTable(&Athlete{}, "Event", &AthleteEvent{})
