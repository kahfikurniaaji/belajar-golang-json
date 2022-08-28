package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("Customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	err = reader.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
