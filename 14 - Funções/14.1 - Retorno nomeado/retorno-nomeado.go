package main

import "fmt"

func calc(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calc(30, 15)
	fmt.Println(soma, sub)

}
