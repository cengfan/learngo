package main

import "fmt"

func main() {

	nextInt :=initSeq()
	fmt.Println(nextInt(),nextInt(),nextInt(),nextInt())
}

func initSeq() func() int  {
	i :=0
	return func() int {
		i +=1
		return i
	}
}
