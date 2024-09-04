package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("I bled on a pivotal stretch")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() { //go routine encapsulada retornando o canal
		for {
			channel <- fmt.Sprintf("Received values: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
