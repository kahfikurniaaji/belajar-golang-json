package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, err := os.Create("CustomerOut.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Kahfi",
		MiddleName: "Kurnia",
		LastName:   "Aji",
	}

	encoder.Encode(customer)
}
