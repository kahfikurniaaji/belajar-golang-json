package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	jsonString := `{"id":"P0001", "name":"Apple Mac Book Pro", "price":20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapDecode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "Apple Mac Book Pro",
		"price": 20000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
