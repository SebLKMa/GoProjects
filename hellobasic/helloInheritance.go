package main

import "fmt"

type Person struct {
    Name string
}

func (pPerson *Person) sayHello() {
    fmt.Println("Hello, my name is ", pPerson.Name)
}

// inheritance via embedded(anonymous) type
// Engineer is-a Person
type Engineer struct {
    Person // anonymous type
    Specialization string
}

func f1() {
    e := new(Engineer)
    e.Person.sayHello()
    e.sayHello()
}

func f2() {
    e := Engineer{Person{"sam"}, "sofware engineering"}
    e.sayHello()
    fmt.Println(e.Specialization);
}

func main() {
    f1();
    f2();
}