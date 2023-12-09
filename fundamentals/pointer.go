package main

/*
	Pointers--> A pointer holds the memory address of a value.
	The type *T is a pointer to a T value. Its zero value is nil.
*/

import "fmt"

func main() {
	var javaScript = "awesome"

	//"&" operator give the memory address of the variable
	memoryAddressVariable := &javaScript
	fmt.Println(memoryAddressVariable)

	//"*" operator give the value of the memory address it pointing at.
	exactedValueFromMemoryAddress := *memoryAddressVariable
	fmt.Println(exactedValueFromMemoryAddress)
}
