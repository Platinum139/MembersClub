package models

import "time"

type Member struct {
	name string
	email string
	registrationDate time.Time
}
