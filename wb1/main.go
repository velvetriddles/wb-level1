package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       uint8
}

// Getter for Name (fN + lN)

func (h *Human) GetName() (string, string) {
	return h.firstName, h.lastName
}

// Getter for Age

func (h *Human) GetAge() uint8 {
	return h.age
}

// Getter for uintPtr 

func uint8Ptr(i uint8) *uint8 {
	return &i
}

// SetterHuman

func (h *Human) SetHuman(firstName, lastName *string, age *uint8) {
	if firstName != nil {
		h.firstName = *firstName
	}

	if lastName != nil {
		h.lastName = *lastName

	}

	if age != nil {
		h.age = *age
	}
}

type Action struct {
	Human
}

func NewAction(firstName, lastName string, age uint8) *Action {
	return &Action{Human: Human{firstName: firstName, lastName: lastName, age: age}}
}

func main() {
	act := NewAction("Petrov", "Maxim", 21)
	act.SetHuman(nil, nil, uint8Ptr(30))
	name, surname := act.GetName()
	age := act.GetAge()
	fmt.Printf(`Your name is %s %s,
	age is %d`, name, surname, age)

}
