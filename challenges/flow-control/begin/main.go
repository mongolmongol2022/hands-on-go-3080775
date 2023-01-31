// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		} else {
			fmt.Println("err is nil")
		}
	}()

	// use os package to read the file specified as a command line argument
	// file := os.Args
	bs, err := os.ReadFile(os.Args[1])
	// /workspaces/hands-on-go-3080775/challenges/data.txt

	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// switch file {
	// case nil:
	// 	panic("something bad happened")
	// default:
	// 	fmt.Println("%s.\n", file)
	// }

	// convert the bytes to a string
	// string(file)
	data := string(bs)

	// initialize a map to store the counts
	count := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	// words := strings.Split(file, " ")
	words := strings.Fields(data)

	// capture the length of the words slice
	// fmt.Println("%d", len(words))
	count["words"] = len(words)
	// spew.Dump(count)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	// for _, word := range words {
	// 	switch word {
	// 	case :
	// 		count["letters"] += 1
	// 	case :
	// 		count["numbers"] += 1
	// 	default:
	// 		count["symbols"] += 1
	// 	}
	// }

	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				count["letters"]++
			} else if unicode.IsNumber(char) {
				count["numbers"]++
			} else {
				count["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(count)
}
