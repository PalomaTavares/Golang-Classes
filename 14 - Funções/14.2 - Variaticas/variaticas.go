package main

import "fmt"

// apenas 1 variatico por funcao e deve ser o ultimo parametro
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func write(text string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(text, numero)
	}
}

func main() {
	totalS := soma(1, 2, 3, 4, 5, 6, 7, 8, 45, 654, 634)
	fmt.Println(totalS)

	write("hello", 11, 13, 56, 78, 97, 98)
}
