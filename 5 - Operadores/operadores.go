package main

import "fmt"

func main() {
	//aritméticos +-/*%
	soma := 2 + 2
	sub := 5 - 2
	div := 14 / 2
	mult := 3 * 3
	restoDiv := 13 % 2

	fmt.Println(soma, sub, div, mult, restoDiv)

	//para fazer operaçẽs as variaveis devem ser do mesmo tipo
	var n1 int16 = 10
	var n2 int16 = 22
	soma2 := n1 + n2

	fmt.Println(soma2)

	////////////////
	//atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	///////////////////////
	//relacionais
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(variavel1 == variavel2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 1)
	fmt.Println(5 != 6)
	////////////////////////
	//logicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	///////////////////
	//unario
	num := 10
	num++
	fmt.Println(num)
	num += 5
	fmt.Println(num)
	num--
	fmt.Println(num)
	num -= 2
	fmt.Println(num)
	num *= 2
	fmt.Println(num)
	num /= 2
	fmt.Println(num)
	num %= 2
	fmt.Println(num)
	/////////////////

}
