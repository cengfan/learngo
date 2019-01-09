package main

import "fmt"

type person struct {
	name string
	age int

}

func main() {
	fmt.Println(person{"zhangsan",11})

	fmt.Println(person{name:"lisi",age:22})


	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp :=&s

	fmt.Println(sp.age)

	sp.age = 30
	fmt.Println(sp.age)
}