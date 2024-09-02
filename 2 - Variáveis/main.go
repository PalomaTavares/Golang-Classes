package main

import "fmt"

func main() {
	//tipo 1
	var variavel1 string = "variavel 1"

	//tipo 2
	variavel2 := "variavel 2" //inferência de tipos
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//tipo 3
	var (
		variavel3 string = "lalalalala"
		variavel4 string = "lalalalalee"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	//tipo 4
	variavel5, variavel6 := "a little bird", "layed down on henry lee"
	fmt.Println(variavel5, variavel6)

	//tipo 5
	const constante1 string = "henry lee - nick cave and the bad seed ft pj harvey"
	fmt.Println(constante1)

	//invertendo variáveis sem aux
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
