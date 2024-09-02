package main

import "fmt"

//metodo é parte de uma estrutura, voce especifica a estrutura a qual está se referindo antes

type user struct {
	nome  string
	idade int8
}

func (u user) salvar() {
	fmt.Printf("Salvando %s\n", u.nome)

}

func (u user) maiorDeIdade() bool { //copia
	return u.idade >= 18

}

func (u *user) fazerAniversario() { //ponteiro
	u.idade++
}

func main() {
	usuario1 := user{"u1", 21}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := user{"u2", 33}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
