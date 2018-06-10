package controllers

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

var people []Person

func init() {
	people = append(people, Person{ID: 1, Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: 2, Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: 3, Firstname: "William", Lastname: "Olojede"})
}



