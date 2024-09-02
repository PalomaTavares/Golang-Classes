package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//go permite mais de um retorno

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(2, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("texxxto f")

	resultadoSoma, resultadoSub := calculos(15, 10)
	fmt.Println(resultadoSoma, resultadoSub)

}
