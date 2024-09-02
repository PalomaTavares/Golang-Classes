package main

import "fmt"

func main() {

	// func(text string) {
	// 	fmt.Println(text)
	// }("parameter")

	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("parameter")

	fmt.Println(retorno)
}
