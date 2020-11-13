package main

import (
	"fmt"
	
)

func main() {
	go printSome()
	fmt.Print("go channels")
	

}

func printSome() {
	fmt.Print("normal func")
	

}