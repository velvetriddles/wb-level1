package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Структура для распаршивания xml данных, потому что в мапу не парсится
type Person struct {
	Name string `xml:"name"`
	Age  string `xml:"age"`
}

// Интерфейс, который ожидает JSON-данные
type JSONParser interface {
	ParseJSON() map[string]string
}

// Реализация JSON-декодера
type JSONData struct {
	data string
}

func (j *JSONData) ParseJSON() map[string]string {
	var result map[string]string
	json.Unmarshal([]byte(j.data), &result)
	return result
}

// Структура для Xml данных
type XMLData struct {
	data string
}

func (x *XMLData) ParseXML() map[string]string {
	// Парсим Xml в структуру 
	var person Person
	xml.Unmarshal([]byte(x.data), &person)

	// Преобразуем структуру в мапу
	result := map[string]string{
		"name": person.Name,
		"age":  person.Age,
	}
	return result
}

// Адаптер для Xml данных, который реализует интерфейс jsonparser
type XMLToJSONAdapter struct {
	xmlData *XMLData
}

func (adapter *XMLToJSONAdapter) ParseJSON() map[string]string {
	// Используем метод ParseXML для получения данных в формате JSON
	return adapter.xmlData.ParseXML()
}

func main() {
	// JSON data
	jsonData := &JSONData{data: `{"name": "Petrovich", "age": "30"}`}
	fmt.Println("JSON Data:", jsonData.ParseJSON())

	// Xml data
	xmlData := &XMLData{data: `<Person><name>Lovetskiy</name><age>18</age></Person>`}

	// Используем адаптер для работы с XML данными как с JSON
	adapter := &XMLToJSONAdapter{xmlData}
	fmt.Println("Adapted Xml Data:", adapter.ParseJSON())
}
