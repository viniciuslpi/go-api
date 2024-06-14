package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Error struct {
	Message string `json:"message,omitempty"`
}

var people []Person

// função principal
func main() {
	// db
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()

	router.HandleFunc("/contato", FindAll).Methods("GET")
	router.HandleFunc("/contato/{id}", FindOne).Methods("GET")
	router.HandleFunc("/contato/{id}", Create).Methods("POST")
	router.HandleFunc("/contato/{id}", Update).Methods("PATCH")
	router.HandleFunc("/contato/{id}", Delete).Methods("DELETE")

	router.HandleFunc("/pokemons", FindAllPokemons).Methods("GET")

	fmt.Println("Running on port: 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
