package main

import (
    "fmt"
    "math/rand"
)

func CalculateValue(values chan int) {
    value := rand.Intn(50)
    fmt.Println("Calculated Random Value: ", value)
    values <- value
}

func main() {
     values := make(chan int)
  defer close(values)

    go CalculateValue(values)

    value := <-values
    fmt.Println(value)
}