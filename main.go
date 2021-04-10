package main

import "fmt"

// this is the go version of OOP
// the Dog struct is a blueprint
type Dog struct {
	name  string
	breed string
	age   int
}

// the printName method is a getter
func (dog Dog) printName() {
	fmt.Println(dog.name)
}

func (dog Dog) printAge() {
	fmt.Println(dog.age)
}

// the birthday method is a setter
func (dog *Dog) birthday() {
	dog.age = dog.age + 1
}

func main() {
	// we create a new dog using the struct
	dog := Dog{
		name:  "tilly",
		age:   7,
		breed: "border collie",
	}
	// we print its name
	dog.printName()
	// we print its age
	dog.printAge()
	// we add 1 to the dogs age
	dog.birthday()
	// we print the age again
	dog.printAge()
}
