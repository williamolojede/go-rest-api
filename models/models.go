package models

type Person struct {
	ID        int   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
}

var People []Person

func init() {
	People = append(People, Person{ID: 1, Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	People = append(People, Person{ID: 2, Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	People = append(People, Person{ID: 3, Firstname: "William", Lastname: "Olojede"})
}