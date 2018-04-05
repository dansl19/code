package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//First we need to declare new struct type
type person struct {
	firstName   string
	lastName    string
	zip         int
	contactInfo // But if you want you can use field contactInfo without contac

}

func main() {
	//alex := person{firstName: "Alex", lastName: "Kubov"} // here we create new var of type person and assign two vars

	//here we declare var "dan" of type person that includes two structure
	dan := person{
		firstName: "Dan", //Here yuo upate the value of var dan
		lastName:  "Slutzky",
		zip:       10000,
		contactInfo: contactInfo{ // here you must to use type and filed the same name
			email:   "dan@gmail.com",
			zipCode: 10409,
		},
	}

	/* but if you want to update value of out structure you will be asked to use pointer
	as "structr" is one of the tupes that requesrs pointers. */

	danPointer := &dan // here we are getting the address in memory of var "dan"

	danPointer.updateName("Denis") // danPointer actually will be used later in the func "updateName" as reciever "p"

	dan.print()
	//fmt.Printf("%+v", dan) // here you print two values with filed of the values
}

// Here we create a function which will print details of person
func (p person) print() {
	fmt.Printf("%+v", p) // here you print two values with filed of the values
}

//this is declaration of new function which has a reciever "p" of type pointer in "person" and this function will excpect for one value "newFirstName"
//"p" == danPointer,  with help of "p" which contains address of "dan" var we updates the value of "dan" var to the values "Denis"
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName // this func will update name of dan for example

}
