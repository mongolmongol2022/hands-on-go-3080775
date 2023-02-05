// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	// var i any = "hello"
	// fmt.Println(i.(string))

	var i any = 1

	// perform a type assertion and handle the error
	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	}

	var j any = "hello"

	// perform a type assertion and handle the error
	if _, ok := j.(string); ok {
		fmt.Printf("%v is a string\n", j)
	}

}
