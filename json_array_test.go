package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		Firstname: "Muhammad",
		Lastname:  "Deril",
		Age:       27,
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","LastName":"Deril","Age":27,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		Firstname: "Muhammad",
		Address: []Address{
			{
				Street:     "Jalan Sesama",
				Country:    "Indonesia",
				PostalCode: "6969",
			},
			{
				Street:     "Jalan Yang Lurus",
				Country:    "Indonesia",
				PostalCode: "72777777",
			},
		},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","LastName":"","Age":0,"Hobbies":null,"Address":[{"Street":"Jalan Sesama","Country":"Indonesia","PostalCode":"6969"},{"Street":"Jalan Yang Lurus","Country":"Konoha","PostalCode":"72777777"}]}`
	jsonByte := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)

	}
	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Address)
	fmt.Println(customer.Address[0])
	fmt.Println(customer.Address[0].Country)
}
