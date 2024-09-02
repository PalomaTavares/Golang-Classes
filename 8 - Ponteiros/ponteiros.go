package main

import "fmt"

func main() {

	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)
	variavel2--
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma ref de memória

	var variavel3 int = 100
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) //retorna endereco de memoria
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro) //retorna o conteudo do endereco de memoria

}
