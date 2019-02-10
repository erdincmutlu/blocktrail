package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID int `json:"id"`
	Verified bool `json:"verified"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Pregnant string `json:"pregnant,omitempty"`
	Child string `json:"child,omitempty"`
	Chilgender string `json:"childgender,omitempty"`
	Childage string `json:"childage,omitempty"`
	Country string `json:"country,omitempty"`
}

var people []Person

func main() {
	router := mux.NewRouter()
	people = append(people, Person{0, true, "name0", "email0", "female", "pregnant_yes", "child_yes", "child_male", "child_5", "Country0"})
	people = append(people, Person{ID:1, Verified:false})
	router.Path("/person").Queries("id", "{id}").HandlerFunc(GetPerson).Methods("GET")
	router.HandleFunc("/addperson", AddPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := r.FormValue("id")
	log.Println("GetPerson:", params, id)

	log.Println("People:", people)

	p, err := ContractGetPerson(params["id"])
	if err == nil {
		json.NewEncoder(w).Encode(p)
	} else {
		notFound := struct {
			ID       string `json:"id"`
			Verified bool   `json:"verified"`
			Error    string `json:"error"`
		}{id, false, "ID not found"}
		json.NewEncoder(w).Encode(notFound)
	}
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Result string `json:"result,omitempty"`
	}

	var person Person
	log.Println("r.Body:", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Verified = true
	log.Println("Addperson:", person)

	err := ContractAddPerson(person)

	var res Result
	if err == nil {
		res = Result{"Ok"}
	} else {
		res = Result{"Not ok"}
		}
	json.NewEncoder(w).Encode(res)
}
