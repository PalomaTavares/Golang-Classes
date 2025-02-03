package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type doggo struct {
	Name  string `json:"name"`
	Breed string `json: "breed"`
	Age   uint   `json: "age"`
}

func main() {
	//json.Marshal()   // struct or method to json
	//json.Unmarshal() //json to struct or method
	dog := doggo{"Niko", "Dálmata", 3}
	dogJSON, error := json.Marshal(dog) // struct or method to json

	if error != nil {
		log.Fatal("error")
	}

	dog2 := map[string]string{
		"name":  "lili",
		"breed": "poodle",
	}

	fmt.Println(dogJSON)
	fmt.Println(bytes.NewBuffer(dogJSON))

	dog2JSON, error := json.Marshal(dog2)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(dog2JSON)
	fmt.Println(bytes.NewBuffer(dog2JSON))

}
