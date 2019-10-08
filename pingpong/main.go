package main

import (
	"fmt"
	"time"
)

func ping(ball chan<- int, action chan<- string) {
	ball <- 1
	action <- "Player Ping"
}

func pong(ball chan<- int, action chan<- string) {
	ball <- 2
	action <- "Player Pong"
}

func referee(action <-chan string) {
	for {
		fmt.Println("Action: ", <-action)
	}
}

func pingPong() {
	ball := make(chan int)
	action := make(chan string)
	go referee(action)
	go ping(ball, action)
	for {
		value := <-ball
		switch value {
		case 1:
			go pong(ball, action)
		case 2:
			go ping(ball, action)
		}
	}
}

func main() {
	go pingPong()
	time.Sleep(10 * time.Second)
}
