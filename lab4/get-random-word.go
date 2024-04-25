package lab4

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
)

type Animal struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetRandomWord() string {

	file, err := ioutil.ReadFile("lab4/animals.json")
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	var animals []Animal
	err = json.Unmarshal(file, &animals)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	return animals[rand.Intn(len(animals))].Name
}
