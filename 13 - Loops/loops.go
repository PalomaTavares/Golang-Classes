package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j += 3 {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"sixxx", "mick", "tom"}
	////////primeiro sempre é o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}
	for indice, letra := range "Smokin' in the boys room" {
		fmt.Println(indice, string(letra))
	}
	mv := map[string]string{
		"band":       "motley crue",
		"song":       "too young to fall in love",
		"other song": "girls girls girls",
	}
	for chave, valor := range mv {
		fmt.Println(chave, valor)
	}
}
