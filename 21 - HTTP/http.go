package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users here"))
}

func main() {
	// HTTP is a communication protocol - it is the base of web communication

	//model CLIENT (MAKES REQUEST) - SERVER (PROCESSES REQUEST AND SENDS RESPONSE)

	//ROUTES - WAY TO IDENTIFY MESSAGE TYPE AND SERVER'S PROCESS

	//URI  - resourse identificator
	//METHOD  - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil)) //server at 5.000

}
