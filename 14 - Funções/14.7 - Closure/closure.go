package main

import "fmt"

func closure() func() {
	txt := "inside closure function"
	funcao := func() {
		fmt.Println(txt)
	}
	return funcao
}

func main() {
	txt := "Dentro da main"
	fmt.Println(txt)

	newFunc := closure() //lembraca de onde ela foi declarada
	newFunc()
}
