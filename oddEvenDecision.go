package main

import "fmt"

func main() {
	var userNo int
	fmt.Print("Enter a number")
	fmt.Scanln(&userNo)
	if userNo%2 == 1 {
		fmt.Println("This number is odd")
	} else if userNo%2 == 0 {
		fmt.Println("This number is even")
	} else {
		fmt.Println("Please enter a number")
	}
}
