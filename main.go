package main

import (
	"fmt"
)

//! learning Goroutines

// func do_something() {
// 	fmt.Println("SLOW : start")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("SLOW : finish")
// }

// func main() {
// 	fmt.Println("MAIN : start")

// 	go do_something()

// 	fmt.Println("MAIN : continue")

// 	time.Sleep(500 * time.Millisecond)

// 	fmt.Println("MAIN : finished")

// 	time.Sleep(1 * time.Second)
// }

//! learning Channels

func main() {
	channel := make(chan string)
	output := make(chan string)

	go func() {
		s := <-channel
		s = s + "GO!"
		output <- s
	}()

	go func() {
		s := "Hello, "
		channel <- s
	}()

	finalString := <-output

	fmt.Println("This is created by channels:", finalString)
}
