package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	//NÃO NECESSARIAMENTE AO MESMO TEMPO REVESAM O TEMPO != DUAS OU MAIS TAREFAS EXECUTADAS CADA UMA EM UM NÚCLEO

	var waitGroup sync.WaitGroup

	waitGroup.Add(6) //how many goroutines

	go func() {
		write("Wake me up when september ends")
		waitGroup.Done()
	}()

	go func() {
		write("summer has come and passed, the innocent can never last")
		waitGroup.Done()
	}()
	go func() {
		write("Wake me up when september eeends")
		waitGroup.Done()
	}()
	go func() {
		write("like my father has come to pass")
		waitGroup.Done()
	}()
	go func() {
		write("seven years has gone so fast")
		waitGroup.Done()
	}()
	go func() {
		write("Wake me up when september ends")
		waitGroup.Done()
	}()

	// go write("Wake me up when september ends") //go routine nâo espera terminar de executar para seguir o programa
	// write("summer has come and passed, the innocent can never last")

	waitGroup.Wait() // -1 routine

}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
