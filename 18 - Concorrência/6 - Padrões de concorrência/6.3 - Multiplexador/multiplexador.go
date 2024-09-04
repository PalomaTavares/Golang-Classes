package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(write("To see if I drew pity"), write("Or pretty litanies from the"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}

	}()
	return canalSaida
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() { //go routine encapsulada retornando o canal
		for {
			channel <- fmt.Sprintf("Received values: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
