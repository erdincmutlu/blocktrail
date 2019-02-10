package main

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

func ContractGetPerson(ID string) (Person, error) {
	for _, item := range people {
		if strconv.Itoa(item.ID) == ID {
			log.Println("item", item)
			return item, nil
		}
	}

	var p Person
	return p, errors.New("ID not found")
}

func ContractAddPerson(person Person) (error) {
	people = append(people, person)
	return nil
}


