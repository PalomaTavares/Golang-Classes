package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())

}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}
func (c circulo) area() float64 {
	return 2 * c.raio * 3.14
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	r := retangulo{10, 25}
	escreverArea(r)
	c := circulo{10}
	escreverArea(c)
}
