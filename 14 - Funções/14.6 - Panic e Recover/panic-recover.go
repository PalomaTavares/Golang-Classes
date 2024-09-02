package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("execução recuperada com sucesso")
	}
}

func estudanteFoiAprovade(n1, n2 float64) bool {
	defer recuperarExecucao() //se nao panic - nil, vai ser ignorado
	med := (n1 + n2) / 2

	if med > 6 {
		return true
	} else if med < 6 {
		return false
	}
	//panic nao e erro
	panic("AAAAAAAAAA A média é EXATAMENTE 6") //para a execucao e chama todos os defer, o programa morre
}

func main() {

	fmt.Println(estudanteFoiAprovade(6, 6))
	fmt.Println("cabô")

}
