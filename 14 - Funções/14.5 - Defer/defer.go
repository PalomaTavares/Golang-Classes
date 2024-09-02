package main

import "fmt"

func f1() {
	fmt.Println("Executando funcao 1")

}
func f2() {
	fmt.Println("Executando funcao 2")

}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("entrou na função alunoAprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Defer")
	defer f1() //adia funcao 1 até o ultimo momento possível
	f2()
	fmt.Println(alunoAprovado(8, 9))

}
