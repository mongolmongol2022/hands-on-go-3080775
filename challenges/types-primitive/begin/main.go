// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "global"

func printoutGlobalVariable() {
	fmt.Printf("%T, global val = %v\n", val, val)
}

func updateGlobalVariable(input string) {
	val = input
}

func main() {
	// create a local variable "val" and assign it an int value
	val := 12345678

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printoutGlobalVariable()

	// update the package-level variable "val" and print it again
	updateGlobalVariable("updated global")
	printoutGlobalVariable()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	var tmp *int = &val
	*tmp = 100
	fmt.Printf("local val = %v\n", val)
	fmt.Printf("%T, local &val = %v\n", &val, &val)

	*(&val) = 200
	fmt.Printf("%T, local &val = %v\n", &val, &val)

}
