package main

/**
变量定义
 */
import "fmt"

var (
	aa = 1
	bb = "bb"
	ss = "ss"
)

func main() {
	fmt.Print("hello go...")
	varibaleZeroValue()
	varibaleInitValue()
	varibaleTypeDef()
	varibaleShortDef()
}

/**
不初始化值
 */
func varibaleZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q \n", a, b)
}

/**
初始化值
 */
func varibaleInitValue() {
	var a, b int = 3, 4

	var s string = "abc"
	fmt.Println(a, b, s)
}

/**
省略类型
 */
func varibaleTypeDef() {
	var a, b = 3, 4

	var s = "abc"
	fmt.Println(a, b, s)
}

/**
不用var
:=只能在函数内使用
 */
func varibaleShortDef() {
	a, b := 3, 4

	s := "abc"

	b = 5
	fmt.Println(a, b, s)
}
