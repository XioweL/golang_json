package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Moch",
		MiddleName: "Ferdian",
		LastName:   "Alvanza",
		Age:        25,
		Married:    false,
		Hobbies:    []string{"Gaming", "Watching", "Coding"},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Moch","MiddleName":"Ferdian","LastName":"Alvanza","Age":25,"Married":false,"Hobbies":["Gaming","Watching","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer) // MAU POINTER DISINI ATAU DI ATAS customer := &Customer JUGA BISA
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Moch",
		Addresses: []Address{
			{
				Street:     "Jalan Negla Atas",
				City:       "Bandung",
				PostalCode: "1234",
			},
			{
				Street:     "Jalan Negla Tengah",
				City:       "Jakarta",
				PostalCode: "4321",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Moch","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Negla Atas","City":"Bandung","PostalCode":"1234"},{"Street":"Jalan Negla Tengah","City":"Jakarta","PostalCode":"4321"}]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

// JIKA HANYA INGIN ENCODE ARRAY, TIDAK PERLU OBJECT ( CUSTOMER )
func TestOnlyJSONArrayComplex(t *testing.T) {
	jsonString := `[{"Street":"Jalan Negla Atas","City":"Bandung","PostalCode":"1234"},{"Street":"Jalan Negla Tengah","City":"Jakarta","PostalCode":"4321"}]`
	jsonBytes := []byte(jsonString)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

}

func TestOnlyJSON(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Negla Atas",
			City:       "Bandung",
			PostalCode: "1234",
		},
		{
			Street:     "Jalan Negla Tengah",
			City:       "Jakarta",
			PostalCode: "4321",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
