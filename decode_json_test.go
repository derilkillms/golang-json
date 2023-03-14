package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	var jsonString = `{"FirstName":"Muhammad", "LastName":"Deril", "Age":27}`
	var jsonData = []byte(jsonString)
	data := &Customer{}
	var err = json.Unmarshal(jsonData, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data.Firstname)
	fmt.Println(data.Lastname)
	fmt.Println(data.Age)

}
