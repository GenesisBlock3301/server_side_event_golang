package main

import "fmt"

func main() {
	go hello()
	fmt.Println("This happens before the hello()")
}

func hello() {
	fmt.Println("It's most likely you will never see this.")
}
