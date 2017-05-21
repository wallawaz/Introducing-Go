package main

import (
	"fmt"
	"time"
)

//continously ping channel
//only allowed to send --> c
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//continously print channel
//only allowed to receive from c
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
