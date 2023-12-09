package main

import (
	"fmt"
)

type contractInfo struct {
	email   string
	zipCode int
}

/*
If property name and type name are same no need to mention the type
explicitly e.g If property name contact is renamed into contractInfo
only contractInfo can be written like this
type person struct {
	firstName, lastName string
	contractInfo
}
instead of this
type person struct {
	firstName, lastName string
	contract contractInfo
}
*/

// Creating a struct called "person"
type person struct {
	firstName, lastName string
	contract            contractInfo
}

func main() {
	//creating a variable called "alex" or type person
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contract: contractInfo{
			email:   "alex@gmail.com",
			zipCode: 123456,
		},
	}
	/* if properties name are omit while assigning a variable then oder of passing
	property value and then order which properties are written in struct have to be equal.

	Alex := person{"Alex","Anderson"}  --> {firstName: Alex, lastName: Anderson}
	Alex := person{"Anderson","Alex"}  --> {firstName: Anderson, lastName: Alex}
	*/
	// update value
	// alex.firstName = "alexander"
	fmt.Println(alex)

	//Struct fields are accessed using a dot.
	fmt.Println(alex.firstName)

	/*If a variable is declare but no value assign to it
	It will set respective zero value for each property
	*/
	var tom person
	fmt.Printf("%+v", tom) // {firstName: '', lastName: ''}

	//Receiver function
	alex.introduction()
	alex.updateName("Alexander", "the Great")

	//use pointer into a normal function
	fmt.Printf("previous email-> %s \n", alex.contract.email)
	updateEmail(&alex, "alexander@gmail.com")
	fmt.Printf("new email-> %s \n", alex.contract.email)
}

func (p person) introduction() {
	fmt.Printf("\nHi, I'm %s %s. Connect me-%s \n", p.firstName, p.lastName, p.contract.email)
}

// Passing the reference using pointer
// The updateName function expects a receiver of type pointer to a person struct
func (p *person) updateName(newFirstName, newLastName string) {
	// fmt.Println(p) --> &{Alex Anderson {alex@gmail.com 123456}}
	(*p).firstName = newFirstName
	(*p).lastName = newLastName
	fmt.Printf("Hi, I just changed my name %s %s\n", p.firstName, p.lastName)
}

func updateEmail(p *person, newEmail string) {
	p.contract.email = newEmail
	fmt.Printf("Hi, I just changed my email\n")
}
