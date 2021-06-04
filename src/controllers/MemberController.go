package controllers

import (
	template2 "html/template"
	"log"
	"net/http"
)

func HandleMembers(w http.ResponseWriter, _ *http.Request) {
	template, err := template2.ParseFiles("src/static/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
