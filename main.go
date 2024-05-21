package main

import (
	"awesomeProject/greet"
	"fmt"
)

func main() {

	var name string
	fmt.Println("Enter your name:")
	scan, err := fmt.Scanf("%s\n", &name)
	if err == nil && scan > 0 {
		text := greet.Greet(name)
		fmt.Println(text)
	} else {
		fmt.Println("Incorrect name")
	}
}
