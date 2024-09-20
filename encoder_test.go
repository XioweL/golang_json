package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writerJson, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writerJson)

	customer := Customer{
		FirstName:  "Moch",
		MiddleName: "Ferdian",
		LastName:   "Alvanza",
		Age:        25,
	}
	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
