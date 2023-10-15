package main

import (
	"encoding/json"
	"log"
)

var sampleJson = `{
	"username": "herr.wayne",
	"password": "L3akyPa5sw0rd!",
	"name": "Herr Wayne",
	"website": "https://example.de",
	"number": 123
}`

func main() {
	rawData := []byte(sampleJson)
	var payload interface{}                  //The interface where we will save the converted JSON data.
	err := json.Unmarshal(rawData, &payload) // Convert JSON data into interface{} type
	if err != nil {
		log.Fatal(err)
	}
	m := payload.(map[string]interface{}) // To use the converted data we will need to convert it into a map[string]interface{}

	log.Printf("Username: %s", m["username"].(string))
	log.Printf("Password: %s", m["password"].(string))
	log.Printf("Name: %s", m["name"].(string))
	log.Printf("Website: %s", m["website"].(string))
	log.Printf("Number: %d", int(m["number"].(float64)))
}
