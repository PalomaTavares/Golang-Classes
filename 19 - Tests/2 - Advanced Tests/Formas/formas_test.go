package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	// TDD - Test Driven Development - fazer os testes antes de fazer a função
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
