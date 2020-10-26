package main

import "fmt"

func main() {

	fmt.Println("pointers")
	var s1 string = "RajyalakshmiUpputuri"
	var s2 *string = &s1
	fmt.Println("s2 Ponting the address of s1 :", s2, "\t string is:", s1)
	fmt.Println("now dereferencing (s2 pointing the value of s1):", *s2)
	/*
		if we set *s2="Upputuri" now s2(derefernce) is going to print "Upputri"
		actually it's also changes s1 to " Upputuri"(string)
		bcz s2 variable is pointing to the s1's address .so,any changes in s2 will refelects in s1 too
	*/
	*s2 = "Upputuri"
	fmt.Println("intialized to a new value for s2:", *s2, "\t Reflecting in s1:", s1)

}
