package schema

import "time"

type AthleteDTO struct {
	ID         RegularIDs `json:"ID"`
	GovID      string     `json:"GovID"`
	FirstNames string     `json:"FirstNames"`
	LastNames  string     `json:"LastNames"`
	PhoneNum   string     `json:"PhoneNumber"`
	Email      string     `json:"Email"`
	Gender     Gender     `json:"Gender"`
	Inscripted bool       `json:"inscripted"`
	Regular    bool       `json:"Regular"`
}

type MajorGetDTO struct {
	ID   RegularIDs `json:"ID"`
	Name string     `json:"Name"`
}

type DisciplineGetBareDTO struct {
	ID   RegularIDs `json:"ID"`
	Name string     `json:"Name"`
}

type UniversityGetBareDTO struct {
	ID    RegularIDs `json:"ID"`
	Name  string     `json:"Name"`
	Local bool       `json:"Local"`
}

type TeamGetBareDTO struct {
	ID         RegularIDs           `json:"ID"`
	Name       string               `json:"Name" binding:"required"`
	Regular    bool                 `json:"Regular" binding:"required"`
	Category   string               `json:"Category"`
	University UniversityGetBareDTO `json:"University"`
	Athletes   []AthleteDTO         `json:"Athletes" binding:"required"`
}

type TeamGetDTO struct {
	ID           RegularIDs `json:"ID"`
	Name         string     `json:"Name" binding:"required"`
	Regular      bool       `json:"Regular" binding:"required"`
	Category     string     `json:"Category"`
	DisciplineID RegularIDs `json:"DisciplineID"`
	UniversityID RegularIDs `json:"UniversityID"`
	Athletes     []Athlete  `json:"Athletes" binding:"required"`
}

type TeamUpdateDTO struct {
	Name       *string   `json:"Name" binding:"omitempty,min=3"`
	Regular    *bool     `json:"Regular"`
	Category   *string   `json:"Category"`
	AthleteIDs []Athlete `json:"AthleteIDs"`
}

type TeamPostDTO struct {
	ID           RegularIDs   `json:"ID"`
	Name         string       `json:"Name" binding:"required"`
	Regular      bool         `json:"Regular" binding:"required"`
	Category     string       `json:"Category"`
	DisciplineID RegularIDs   `json:"DisciplineID"`
	UniversityID RegularIDs   `json:"UniversityID"`
	AthleteIDs   []RegularIDs `json:"AthleteIDs" binding:"required"`
}

type TeacherGetBareDTO struct {
	ID          RegularIDs             `json:"ID"`
	FirstNames  string                 `json:"FirstNames"`
	LastNames   string                 `json:"LastNames"`
	GovID       string                 `json:"GovID"`
	Disciplines []DisciplineGetBareDTO `json:"Disciplines"`
}

type TeacherGetDTO struct {
	ID          RegularIDs             `json:"ID"`
	FirstNames  string                 `json:"FirstNames"`
	LastNames   string                 `json:"LastNames"`
	PhoneNum    string                 `json:"PhoneNumber"`
	Email       string                 `json:"Email"`
	GovID       string                 `json:"GovID"`
	Disciplines []DisciplineGetBareDTO `json:"Disciplines"`
	// Events      []Event      `json:"events"`
}
type TeacherCreateDTO struct {
	FirstNames    string       `json:"FirstNames" binding:"required,min=2"`
	LastNames     string       `json:"LastNames" binding:"required,min=2"`
	PhoneNum      string       `json:"PhoneNumber"`
	Email         string       `json:"Email" binding:"omitempty,email"`
	GovID         string       `json:"GovID" binding:"required,len=8"`
	DisciplineIDs []RegularIDs `json:"DisciplineIDs"`
}

type TourneyGetBareDTO struct {
	ID     RegularIDs `json:"ID"`
	Name   string     `json:"Name"`
	Status Status     `json:"Status"`
}

type TourneyGetFullDTO struct {
	ID     RegularIDs        `json:"ID"`
	Name   string            `json:"Name"`
	Status Status            `json:"Status"`
	Events []EventGetBareDTO `json:"Events"`
}

type EventGetBareDTO struct {
	ID                 RegularIDs           `json:"ID"`
	Name               string               `json:"Name"`
	Date               time.Time            `json:"Date"`
	Status             string               `json:"Status"`
	HomePoints         uint8                `json:"HomePoints"`
	OppositePoints     uint8                `json:"OppositePoints"`
	HomeTeam           TeamGetBareDTO       `json:"HomeTeam"`
	OppositeTeam       TeamGetBareDTO       `json:"OppositeTeam"`
	ResponsableTeacher TeacherGetBareDTO    `json:"ResponsableTeacher"`
	Discipline         DisciplineGetBareDTO `json:"Discipline"`
}

type EventGetDTO struct {
	ID                 RegularIDs           `json:"ID"`
	Name               string               `json:"Name"`
	Date               time.Time            `json:"Date"`
	Status             string               `json:"Status"`
	Observation        string               `json:"Observation"`
	Ubication          string               `json:"Ubication"`
	HomePoints         uint8                `json:"HomePoints"`
	OppositePoints     uint8                `json:"OppositePoints"`
	HomeTeam           TeamGetBareDTO       `json:"HomeTeam"`
	OppositeTeam       TeamGetBareDTO       `json:"OppositeTeam"`
	Tourney            TourneyGetBareDTO    `json:"Tourney"`
	ResponsableTeacher TeacherGetBareDTO    `json:"ResponsableTeacher"`
	Discipline         DisciplineGetBareDTO `json:"Discipline"`
}
