package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	//NÃO NECESSARIAMENTE AO MESMO TEMPO REVESAM O TEMPO != DUAS OU MAIS TAREFAS EXECUTADAS CADA UMA EM UM NÚCLEO
	go infinityWritting("Wake me up when september ends")
	infinityWritting("summer has come and passed, the innocent can never last")

}

func infinityWritting(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
