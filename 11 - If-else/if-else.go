package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	n := -4

	if n > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if n2 := n; n2 > 0 {
		fmt.Println("numero é maior que 0")
	} else if n < -10 {
		fmt.Println("numero menor que -10")
	} else {
		fmt.Println("está entre 0 e - 10")
	}
}
