package models

import "time"

type Member struct {
	Name string
	Email string
	RegistrationDate time.Time
}
