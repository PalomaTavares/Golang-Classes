package main

import "fmt"

type user struct {
	nome     string
	idade    uint16
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {

	fmt.Println("Arquivo structs")

	var u user
	u.nome = "lestat de lioncourt"
	u.idade = 490

	fmt.Println(u.nome, u.idade)

	enderecoDoLuis := endereco{"red light district - new orleans", 666}

	u2 := user{"luis depont du lac", 200, enderecoDoLuis}
	fmt.Println(u2)

	u3 := user{idade: 120}
	fmt.Println(u3)
}
