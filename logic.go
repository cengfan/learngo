package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println(" 7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println("  小于0")
	} else if num > 0 {
		fmt.Println("大于0")
	}
}
