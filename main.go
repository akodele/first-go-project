package main

import (
	"awesomeProject/greet"
	"fmt"
)

func main() {

	var name string
	fmt.Println("Enter your name:")
	scan1, err1 := fmt.Scanf("%s\n", &name)
	var table int
	fmt.Println("Enter your table:")
	scan2, err2 := fmt.Scanf("%d\n", &table)
	if err1 == nil && scan1 > 0 && err2 == nil && scan2 > 0 {
		text := greet.Greet(name, table)
		fmt.Println(text)
	} else {
		fmt.Println(err1)
		fmt.Println(err2)
	}
}
