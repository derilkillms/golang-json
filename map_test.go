package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSONnya dynamic, artinya attribute nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap. Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct kita harus menentukan semua attributenya.

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Iphone 14 Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P0001",
		"name":      "Iphone 14 Pro",
		"image_url": "http://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
