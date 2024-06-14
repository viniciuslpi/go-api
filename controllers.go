package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			error := Error{"The contact already exists on db"}
			json.NewEncoder(w).Encode(error)
			return
		}
	}

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedPerson Person
	_ = json.NewDecoder(r.Body).Decode(&updatedPerson)

	for index, person := range people {
		if person.ID == params["id"] {
			people[index].Firstname = updatedPerson.Firstname
			people[index].Lastname = updatedPerson.Lastname
			people[index].Address = updatedPerson.Address
			json.NewEncoder(w).Encode(people[index])
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			break
		}
	}

	error := Error{"Contact not found"}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}
