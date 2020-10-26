package main

import "fmt"

func main() {
	fmt.Println("sjdkj")
	/*
	   array which is used to store
	   a same data type  .Below is the very basic example
	   for creating an array
	*/
	var arr [5]string
	arr[0] = "Try"
	arr[1] = "make"
	arr[2] = "it"
	arr[3] = "possible"
	for x := 0; x < 4; x++ {
		fmt.Println(arr[x])
	}
	//here we created an array using shor hand notation
	arr1 := [5]int{10, 20, 30, 40}
	for x := 0; x < 4; x++ {
		fmt.Println(arr1[x])
	}
	/*here array is creat using "ellipses"
	instead of specifing the array length compiler identifies the lenth
	of an array based on the elements we specified
	*/
	arr2 := [...]float32{23.2, 25.4, 12.5, 124.0, 54.6}
	for x := 0; x < len(arr2)-1; x++ {
		fmt.Println(arr2[x])
	}

}
