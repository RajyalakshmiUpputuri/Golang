package main

import (
	"fmt"
	"time"
)

func main() {
	go printSomething()
	fmt.Print("go channels")
	time.Sleep(10 * time.Second)

}

func printSomething() {
	fmt.Print("normal func")
	time.Sleep(1 * time.Millisecond)

}