package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src/static"))
	http.Handle("/", fs)

	log.Println("Starting server...")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
