package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 22
    return &p
}

func main() {

    fmt.Println(person{"Raji", 20})

    fmt.Println(person{name: "Ramma", age: 30})

    fmt.Println(person{name: "Aravi"})

    fmt.Println(&person{name: "Annad", age: 26})

    fmt.Println(newPerson("Jaya"))

    s := person{name: "Puji", age: 18}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}