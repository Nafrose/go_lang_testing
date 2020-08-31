package clihelper

import "time"

type PersonType struct {
	Id          int       `json:"ID"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Alias       string    `json:"alias"`
	DateOfBirth time.Time `json:"dob"`
	Profession  string    `json:"profession"`
	Parent      string    `json:"parent"`
}
