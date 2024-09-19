package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "002",
		Name:     "ASUS RTX 3050",
		ImageURL: "http://www.google.com/rtx.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"002","name":"ASUS RTX 3050","image_url":"http://www.google.com/rtx.png"}`
	jsonBytes := []byte(jsonString)

	product := Product{}

	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
