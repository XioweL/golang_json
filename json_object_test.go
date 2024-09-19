package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Moch",
		MiddleName: "Ferdian",
		LastName:   "ALvanza",
		Age:        25,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
