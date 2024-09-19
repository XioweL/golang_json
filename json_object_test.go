package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, City, PostalCode string
}
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
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
