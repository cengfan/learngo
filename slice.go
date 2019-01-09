package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//s[3] = "d"
	fmt.Println(s)

	fmt.Println(len(s))

	s = append(append(s, "d"), "e")

	fmt.Println(s)

	cps :=make([]string,len(s))

	copy(cps,s)
	fmt.Println("cps",cps)

	fmt.Println(s)


	s1 :=s[2:5]
	fmt.Println("s1",s1)

	s2 :=s[2:]
	fmt.Println("s2", s2)

	s3 :=s[:5]
	fmt.Println("s3",s3)

	t :=[]string{"1","2","3"}
	fmt.Println("t",t)



}
