package main

import (
	"MemberClub/src/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.HandleMembers)
	fs := http.FileServer(http.Dir("src/static"))

	http.Handle("/src/static/", http.StripPrefix("/src/static/", fs))

	log.Println("Starting server...")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
