package main

import (
	"fmt"
	"time"
)

//! learning Goroutines

func do_something() {
	fmt.Println("SLOW : start")
	time.Sleep(1 * time.Second)
	fmt.Println("SLOW : finish")
}

func main() {
	fmt.Println("MAIN : start")

	go do_something()

	fmt.Println("MAIN : continue")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("MAIN : finished")

	time.Sleep(1 * time.Second)
}
