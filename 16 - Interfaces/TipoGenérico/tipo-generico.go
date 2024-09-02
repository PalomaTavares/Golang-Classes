package main

import "fmt"

func generica(interf interface{}) { //aceita qualquer tipo
	fmt.Println(interf) //println é tipo genérico, aceit qualquer tipo
}

func main() {
	generica("str")
	generica(1)
	generica(false)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}
	fmt.Println(mapa)
}
