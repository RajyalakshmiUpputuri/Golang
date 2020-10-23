package main

import "fmt"

func main() {
	var num = 23 //dropping int
	fmt.Println("hello Raji", num)

	var num1 int32 = 72 //adding int
	fmt.Println(num1)

	var a, a1 = 45, 64 //mutip var declaration with dropping int
	fmt.Println(a * a1)

	val, val1 := 34, 23 //short hand declaration
	fmt.Println(val + val1)

	var (
		name = "Raji"
		age  = 20
	)
	fmt.Println(name, "welcome", age)
}
