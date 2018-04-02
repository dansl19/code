package main

import (
	"fmt"
)

//First we need to declare new struct type
type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Kubov"} // here we create new var of type person and assign two vars

	var dan person //there is another way to declare struc by using thee following 3 lines :
	dan.firstName = "Dan"
	dan.lastName = "Slutzky"
	//fmt.Println(alex, dan)

	fmt.Printf("%d - %s\n", alex, dan) // here you print two values with filed of the values
}
