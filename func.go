package main

import "fmt"

func main() {

	c := plus(1, 2)
	fmt.Println("c=", c)
	c1, c2 := vals()
	fmt.Println("c1=", c1, ",c2=", c2)

	_, c3 := vals()
	fmt.Println(c3)

	sum(1, 2, 3)

	a := []int{1, 2, 3}
	sum(a...)
}

/**
单个返回值
 */
func plus(a int, b int) int {

	return a + b
}

/*
多个返回值
 */
func vals() (a int, b int) {
	return 100, 200
}

/**
变长参数
 */
func sum(nums ... int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
