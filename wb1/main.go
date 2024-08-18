package main

import "fmt"

/*
	=== Задача №1 ===

	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской
	структуры Human (аналог наследования).

*/

// Структура Human со свойствами firstName, lastName строковый тип, age - uint8, беззнаковый 8-битный
type Human struct {
	firstName string
	lastName  string
	age       uint8
}

// Метод получения имени, возвращающий два значения (firstName + lastName сразу)
func (h *Human) GetName() (string, string) {
	return h.firstName, h.lastName
}

// Метод получения возраста
func (h *Human) GetAge() uint8 {
	return h.age
}

// Частичное обновление
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

// Структура Action с встраиванием Human
type Action struct {
	Human
}

// Конструктор для создания нового Action
func NewAction(firstName, lastName string, age uint8) *Action {
	return &Action{Human: Human{firstName: firstName, lastName: lastName, age: age}}
}

// Вспомогательная функция для получения указателя на uint8
func uint8Ptr(i uint8) *uint8 {
	return &i
}

// Вспомогательная функция для получения указателя на string
func stringPtr(s string) *string {
	return &s
}

func main() {
	action := NewAction("Petrov", "Maxim", 21)

	// Получаем и выводим начальные данные
	firstName, lastName := action.GetName()
	age := action.GetAge()
	fmt.Printf("Initial state: Your name is %s %s, Age is %d\n", firstName, lastName, age)

	// Изменяем структуру (частичное обновление) передавая нам нужные значения
	action.SetHuman(nil, stringPtr("LovetskiyOneLove"), uint8Ptr(30))

	// Получаем и выводим обновленные данные
	age = action.GetAge()
	firstName, lastName = action.GetName()
	fmt.Printf("After update: Your name is %s %s, Age is %d\n", firstName, lastName, age)
}
