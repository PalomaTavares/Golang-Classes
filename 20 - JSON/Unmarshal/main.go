package main

import (
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
	dogJSON := `{"name":"Niko","Breed":"Dálmata","Age":3}`

	var dog doggo

	if erro := json.Unmarshal([]byte(dogJSON), &dog); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dog)

	dog2JSON := `{"name":"lili", "breed":"poodle"}`

	dog2 := make(map[string]string) //carefull with type

	if erro := json.Unmarshal([]byte(dog2JSON), &dog2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dog2)
}
