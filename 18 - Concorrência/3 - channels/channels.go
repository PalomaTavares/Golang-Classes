package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Do you have the time to listen to me whine", channel)

	fmt.Println("About nothing and everything all at once?")
	// for { //deadlock
	// 	message := <-channel
	// 	fmt.Println(message)
	// }

	for {
		message, open := <-channel
		//is channel still open
		if !open {
			break
		}
		fmt.Println(message)
	}

	for message := range channel { // each received message
		fmt.Println(message)
	}

	fmt.Println("I am one one those melodramatic fools")

}
func write(text string, channel chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		channel <- text //value goes inside channel
		time.Sleep(time.Second)
	}
	close(channel)
}
