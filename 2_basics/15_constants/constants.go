package main

import "fmt"

const pi = 3.14

func main() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
