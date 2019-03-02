package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
