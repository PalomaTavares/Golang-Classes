package main

import "fmt"

var n int

func init() {
	fmt.Println("init é executada antes do main, é uma por arquivo")
	n = 10
}

func main() {
	fmt.Println("MAin")
	fmt.Println(n)
}
