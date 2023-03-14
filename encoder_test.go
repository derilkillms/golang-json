package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, err := os.Create("CustomerOut.json")

	encoder := json.NewEncoder(writer)
	customer := Customer{
		Firstname: "Muhammad",
		Lastname:  "Deril",
		Age:       27,
		Hobbies:   []string{},
		Address:   []Address{},
	}
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}
