package main

import "fmt"

// a function is a block of code which is used to perform a specific task

func sum(a, b int) int { //function declaration
	return a + b
}
func mul(a, b int32) int32 {
	var ans = a * b
	return ans
}

func main() {
	val := sum(12, 13)

	// here we'r discarding parameter using _(blank Identifier)
	fmt.Printf("the sum is :%v", val)
	fmt.Println("\n the mul is", mul(45, 23))

}
