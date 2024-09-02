package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    int8
}
type estudante struct {
	pessoa    //sem tipo para usar os atributos de pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Andrew", "Hozier", 30, 2}
	fmt.Println(p1)

	e1 := estudante{p1, "letras", "ufv"}
	fmt.Println(e1)

}
