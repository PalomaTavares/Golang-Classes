package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")
	//		posicoes	tipo
	var array1 [5]int
	array1[0] = 0
	array1[1] = 1

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array1)
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	fmt.Println(reflect.TypeOf(array3))

	/////////////////////////aqui comeca slice//////////////////////
	//basicamente um array dinâmico
	slice := []int{12, 13, 36, 68, 97, 96, 80, 43, 59}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)

	fmt.Println(slice)

	//inicio é inclusivo, fim é exclusivo
	slice2 := array2[1:4]
	fmt.Println(slice2)

	array2[1] = "posicao alterada"
	fmt.Println(slice2)
	/////////////////////////aqui comecam arrays internos /////////////
	////tipo, tamanho, capacidade
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) //capacidade será aumentada
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(cap(slice4))
}
