package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "002",
		"name":  "ASUS RTX",
		"price": 5000000,
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"002","name":"ASUS RTX 3050","price":5000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

}
