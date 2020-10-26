package main

import "fmt"

func main() {
	fmt.Println("advance")
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	slice := array[1:6]
	fmt.Println("slice ele is:", slice)
	fmt.Printf(" capacity of slice is:%d", cap(slice))
	fmt.Printf(" \nlength of slice is:%d", len(slice))

}
