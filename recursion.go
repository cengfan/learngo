package main

import "fmt"

func main() {
	fmt.Println(face(7))
}

func face(a int) int  {
	if a == 0{
		return 1
	}
	return a * face(a-1)
}