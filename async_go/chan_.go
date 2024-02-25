package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Worker:", i)
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go worker(ch)

	for i := 0; i < 5; i++ {
		val := <-ch
		time.Sleep(2 * time.Second)
		fmt.Println("Main:", val)
	}
}
