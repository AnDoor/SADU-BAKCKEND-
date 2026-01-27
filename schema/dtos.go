package schema

import "time"

type AthleteDTO struct {
	ID         RegularIDs `json:"id"`
	GovID      string     `json:"id_personal"`
	FirstNames string     `json:"name"`
	LastNames  string     `json:"lastname"`
	PhoneNum   string     `json:"phonenumber"`
	Email      string     `json:"email"`
	Inscripted bool       `json:"inscripted"`
	Regular    bool       `json:"regular"`
}
type MajorGetDTO struct {
	ID   RegularIDs `json:"id"`
	Name string     `json:"name"`
}

type DisciplineGetBareDTO struct {
	ID   RegularIDs `json:"id"`
	Name string     `json:"name"`
}

type UniversityGetBareDTO struct {
	ID    RegularIDs `json:"id"`
	Name  string     `json:"name"`
	Local bool       `json:"local"`
}

type TeamGetBareDTO struct {
	ID         RegularIDs           `json:"id"`
	Name       string               `json:"name" binding:"required"`
	Regular    bool                 `json:"regular" binding:"required"`
	Category   string               `json:"category"`
	University UniversityGetBareDTO `json:"university"`
	Athletes   []Athlete            `json:"athletes" binding:"required"`
}

type TeamGetDTO struct {
	ID           RegularIDs `json:"id"`
	Name         string     `json:"name" binding:"required"`
	Regular      bool       `json:"regular" binding:"required"`
	Category     string     `json:"category"`
	DisciplineID RegularIDs `json:"discipline_id"`
	UniversityID RegularIDs `json:"university_id"`
	Athletes     []Athlete  `json:"athletes" binding:"required"`
}

type TeamUpdateDTO struct {
	Name       *string   `json:"name" binding:"omitempty,min=3"`
	Regular    *bool     `json:"regular"`
	Category   *string   `json:"category"`
	AthleteIDs []Athlete `json:"athlete_ids"`
}

type TeamPostDTO struct {
	ID           RegularIDs   `json:"id"`
	Name         string       `json:"name" binding:"required"`
	Regular      bool         `json:"regular" binding:"required"`
	Category     string       `json:"category"`
	DisciplineID RegularIDs   `json:"discipline_id"`
	UniversityID RegularIDs   `json:"university_id"`
	AthleteIDs   []RegularIDs `json:"athlete_ids" binding:"required"`
}

type TeacherGetBareDTO struct {
	ID          RegularIDs             `json:"id"`
	FirstNames  string                 `json:"first_names"`
	LastNames   string                 `json:"last_names"`
	GovID       string                 `json:"gov_id"`
	Disciplines []DisciplineGetBareDTO `json:"disciplines"`
}

type TeacherGetDTO struct {
	ID          RegularIDs             `json:"id"`
	FirstNames  string                 `json:"first_names"`
	LastNames   string                 `json:"last_names"`
	PhoneNum    string                 `json:"phone_num"`
	Email       string                 `json:"email"`
	GovID       string                 `json:"gov_id"`
	Disciplines []DisciplineGetBareDTO `json:"disciplines"`
	// Events      []Event      `json:"events"`
}
type TeacherCreateDTO struct {
	FirstNames    string       `json:"first_names" binding:"required,min=2"`
	LastNames     string       `json:"last_names" binding:"required,min=2"`
	PhoneNum      string       `json:"phone_num"`
	Email         string       `json:"email" binding:"omitempty,email"`
	GovID         string       `json:"gov_id" binding:"required,len=8"`
	DisciplineIDs []RegularIDs `json:"discipline_ids"`
}

type TourneyGetBareDTO struct {
	ID     RegularIDs `json:"id"`
	Name   string     `json:"name"`
	Status Status     `json:"status"`
}

type TourneyGetFullDTO struct {
	ID     RegularIDs        `json:"id"`
	Name   string            `json:"name"`
	Status Status            `json:"status"`
	Events []EventGetBareDTO `json:"events"`
}

type EventGetBareDTO struct {
	ID                 RegularIDs           `json:"id"`
	Name               string               `json:"name"`
	Date               time.Time            `json:"date"`
	Status             string               `json:"status"`
	HomePoints         uint8                `json:"home_points"`
	OppositePoints     uint8                `json:"opposite_points"`
	HomeTeam           TeamGetBareDTO       `json:"home_team"`
	OppositeTeam       TeamGetBareDTO       `json:"opposite_team"`
	ResponsableTeacher TeacherGetBareDTO    `json:"responsable_teacher"`
	Discipline         DisciplineGetBareDTO `json:"discipline"`
}

type EventGetDTO struct {
	ID                 RegularIDs           `json:"id"`
	Name               string               `json:"name"`
	Date               time.Time            `json:"date"`
	Status             string               `json:"status"`
	Observation        string               `json:"observation"`
	Ubication          string               `json:"ubication"`
	HomePoints         uint8                `json:"home_points"`
	OppositePoints     uint8                `json:"opposite_points"`
	HomeTeam           TeamGetBareDTO       `json:"home_team"`
	OppositeTeam       TeamGetBareDTO       `json:"opposite_team"`
	Tourney            TourneyGetBareDTO    `json:"tourney"`
	ResponsableTeacher TeacherGetBareDTO    `json:"responsable_teacher"`
	Discipline         DisciplineGetBareDTO `json:"discipline"`
}
