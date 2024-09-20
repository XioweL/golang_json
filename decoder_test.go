package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {

	readJson, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(readJson)

	customer := Customer{}
	decoder.Decode(&customer)

	fmt.Println(customer)

}
