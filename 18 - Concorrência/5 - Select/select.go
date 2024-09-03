package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			c1 <- "channel 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "channel 2"
		}
	}()

	for {
		select { //if c1 is ready prints c1 if c2 is ready print c2
		case messageC1 := <-c1:
			fmt.Println(messageC1)
		case messageC2 := <-c2:
			fmt.Println(messageC2)
		}
	}
}
