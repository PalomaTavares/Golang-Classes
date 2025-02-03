package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())

}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}
