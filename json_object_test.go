package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
