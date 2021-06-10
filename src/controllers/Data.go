package controllers

import (
	"MemberClub/src/models"
	"regexp"
	"time"
)

type Data struct {
	Members []models.Member
	Errors map[string]string
}

func (data *Data) Validate(name string, email string) bool {
	validateName := func(name string) bool {
		rgx := regexp.MustCompile("^([A-Za-z]+(\\s[a-zA-Z]+)*)$")
		return rgx.Match([]byte(name))
	}
	validateEmail := func(email string) bool {
		rgx := regexp.MustCompile(
			"^([A-Za-z0-9_\\.]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+(\\.[A-Za-z0-9]+)*)$")
		return rgx.Match([]byte(email))
	}
	data.Errors = make(map[string]string)
	if !validateName(name) {
		data.Errors["name"] = "Name can contain only letters."
	}
	if !validateEmail(email) {
		data.Errors["email"] = "Enter valid email, please."
	}
	return len(data.Errors) == 0
}

func (data *Data) MemberExists(email string) bool {
	for _, member := range data.Members {
		if member.Email == email {
			data.Errors["emailExists"] = "Member with this email already exists."
			return true
		}
	}
	return false
}

func (data *Data) AddMember(name string, email string) {
	data.Members = append(data.Members, models.Member{
		Name:             name,
		Email:            email,
		RegistrationDate: time.Now(),
	})
}

