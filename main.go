package main

import "fmt"

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

// func main() {
// 	channel := make(chan string)
// 	output := make(chan string)

// 	go func() {
// 		s := <-channel
// 		s = s + "GO!"
// 		output <- s
// 	}()

// 	go func() {
// 		s := "Hello, "
// 		channel <- s
// 	}()

// 	finalString := <-output

// 	fmt.Println("This is created by channels:", finalString)
// }

//! learning Go's type system

// type Describer interface {
// 	describe() string
// }

// func printDescribe(d Describer) {
// 	fmt.Println(d.describe())
// }

// type Person struct {
// 	first_name string
// 	last_name  string
// 	age        int
// }

// func (p Person) name() string {
// 	return p.first_name + " " + p.last_name
// }

// func (p Person) describe() string {
// 	return fmt.Sprintf("%s is %d years old", p.name(), p.age)
// }

// func main() {
// 	mmd := Person{
// 		first_name: "Mohammad",
// 		last_name:  "Jamali",
// 		age:        19,
// 	}

// 	printDescribe(mmd)
// }

//! learning Go's generics

func sumNumbers[N int | float64](numbers ...N) N {
	var total N
	for i := range numbers {
		total += numbers[i]
	}

	return total
}

func main() {
	fmt.Println(sumNumbers[int](1, 5, 6, 8))
	fmt.Println(sumNumbers(1.5, 5.2, 7.2, 8.0))
}
