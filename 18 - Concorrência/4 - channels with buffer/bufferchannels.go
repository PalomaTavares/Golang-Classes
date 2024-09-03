package main

import "fmt"

func main() {
	channel := make(chan string, 3) //channel with buffer only gets deadlocks if i pass more than 3 values
	channel <- "Do you have the time to listen to me whine"
	channel <- "About nothing and everything all at once?"
	channel <- "I am one of those melodramatic fools"

	message := <-channel
	fmt.Println(message)
}
