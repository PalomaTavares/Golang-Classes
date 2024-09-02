package main

import (
	"errors"
	"fmt"
)

func main() {
	//inteiros: int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)

	var num int = 100000000 //arquitetura do pc como base 64 bits = int64
	fmt.Println(num)

	//uint é para negativos

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE	BYTE = UINT8
	var numero3 rune = 1024
	fmt.Println(numero3)

	//FLOAT: float32 e float64
	var numero4 float32 = 123.45
	fmt.Println(numero4)
	var numero5 float64 = 1230000000.45
	fmt.Println(numero5)

	char := 'B'
	fmt.Println(char)

	//error é um tipo
	var algoErrado error
	fmt.Println(algoErrado) //<nil>

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
