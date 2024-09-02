package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Aplicação de linha de comando")

	application := app.Generate()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
