package controllers

import (
	"MemberClub/src/models"
	"testing"
	"time"
)

// test validate
func TestValidateNameEmailValid(t *testing.T) {
	var data Data
	name := "John"					// valid
	email := "johnny@gmail.com"		// valid
	if data.Validate(name, email) == false {
		t.Errorf("Got: %t\n", false)
		t.Errorf("Expected: %t\n", true)
	}
}

func TestValidateNameEmailInvalid(t *testing.T) {
	var data Data
	name := "John123"				// invalid
	email := "johnny@gmail"			// invalid
	if data.Validate(name, email) == true {
		t.Errorf("Got: %t\n", true)
		t.Errorf("Expected: %t\n", false)
	}
}

func TestValidateNameValidEmailInvalid(t *testing.T) {
	var data Data
	name := "Alice"							// valid
	email := "alice&%thomas@gmail.com"		// invalid
	if data.Validate(name, email) == true {
		t.Errorf("Got: %t\n", true)
		t.Errorf("Expected: %t\n", false)
	}
}

func TestValidateNameInvalidEmailValid(t *testing.T) {
	var data Data
	name := "Alice_Watson"					// invalid
	email := "alice_thomas@gmail.com"		// valid
	if data.Validate(name, email) == true {
		t.Errorf("Got: %t\n", true)
		t.Errorf("Expected: %t\n", false)
	}
}

// test MemberExists
func TestMemberExists_NoMember(t *testing.T) {
	var data Data
	if data.MemberExists("jonny@gmail.com") == true {
		t.Errorf("Got: %t\n", true)
		t.Errorf("Expected: %t\n", false)
	}
}

func TestMemberExists_WithMemberTrue(t *testing.T) {
	data := Data{
		Members: make([]models.Member, 3),
		Errors: make(map[string]string),
	}
	data.Members = append(data.Members, models.Member{
		Name:  "John",
		Email: "johny@gmail.com",
	})
	data.Members = append(data.Members, models.Member{
		Name:  "Alice",
		Email: "alice_thomas@gmail.com",
	})
	data.Members = append(data.Members, models.Member{
		Name:  "Nick",
		Email: "nick.gonzalez@gmail.com",
	})
	if data.MemberExists("alice_thomas@gmail.com") == false {
		t.Errorf("Got: %t\n", false)
		t.Errorf("Expected: %t\n", true)
	}
}

func TestMemberExists_WithMemberFalse(t *testing.T) {
	data := Data{
		Members: make([]models.Member, 3),
		Errors: make(map[string]string),
	}
	data.Members = append(data.Members, models.Member{
		Name:  "John",
		Email: "johny@gmail.com",
	})
	data.Members = append(data.Members, models.Member{
		Name:  "Alice",
		Email: "alice_thomas@gmail.com",
	})
	data.Members = append(data.Members, models.Member{
		Name:  "Nick",
		Email: "nick.gonzalez@gmail.com",
	})
	if data.MemberExists("lia.jerison@gmail.com") == true {
		t.Errorf("Got: %t\n", true)
		t.Errorf("Expected: %t\n", false)
	}
}

// test AddMember
func TestAddMember(t *testing.T) {
	data := Data{
		Members: make([]models.Member, 3),
		Errors: make(map[string]string),
	}
	data.AddMember("Lia Jerison", "lia.jerison@gmail.com")
	regDate := time.Now()

	memberAdded := false
	for _, member := range data.Members {
		if member.Name == "Lia Jerison" && member.Email == "lia.jerison@gmail.com" {
			memberAdded = true
			if member.RegistrationDate.Year() != regDate.Year() ||
				member.RegistrationDate.Month() != regDate.Month() ||
				member.RegistrationDate.Day() != regDate.Day() {
				t.Errorf("Got: RegistratioDate %s\n", member.RegistrationDate.Format("02-01-2006"))
				t.Errorf("Expected: RegistratioDate %s\n", regDate.Format("02-01-2006"))
			}
		}
	}
	if !memberAdded {
		t.Errorf("Got: no member added\n")
		t.Errorf("Expected: member added\n")
	}
}