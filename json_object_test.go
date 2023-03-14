package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname, Lastname string
	Age                 int
	Hobbies             []string
}

func TestJSONOBject(t *testing.T) {
	customer := Customer{
		Firstname: "Muhammad",
		Lastname:  "Deril",
		Age:       27,
	}
	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}
