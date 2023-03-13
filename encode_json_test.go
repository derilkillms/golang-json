package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	logJson("Deril")
	logJson(27)
	logJson(false)
	logJson([]int16{1, 2, 3, 4, 5})
	logJson([]string{"Muhammad", "Deril"})
}
