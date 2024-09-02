package main

import "fmt"

func invertSignal(n int) int {
	return n * -1
}

func invertSignalPointer(num *int) {
	*num = *num * -1
}

func main() {
	n := 20

	invN := invertSignal(n)
	fmt.Println(invN)
	fmt.Println(n) //n nao foi alterada

	newN := 66
	fmt.Println(newN)

	invertSignalPointer(&newN)
	fmt.Println(newN)

}
