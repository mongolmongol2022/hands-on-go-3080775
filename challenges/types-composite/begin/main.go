// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthorName(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	lib := library{}

	// book1 := book{
	// 	title: "title1",
	// 	author: author{
	// 		name: "Tom",
	// 	},
	// }

	// book2 := book{
	// 	title: "title2",
	// 	author: author{
	// 		name: "Jerry",
	// 	},
	// }

	jb := author{
		name: "James baldwin",
	}

	// add 2 books to the library by the same author
	// lib.addBook(book1)
	// lib.addBook(book1)

	lib.addBook(book{
		title:  "Inline1",
		author: jb,
	})

	lib.addBook(book{
		title:  "Inline2",
		author: jb,
	})

	// add 1 book to the library by a different author
	//
	// lib.addBook(book2)

	lib.addBook(book{
		title:  "Inline3",
		author: author{name: "Inline"},
	})

	// dump the library with spew
	// fmt.Printf("%v\n", lib)
	// spew.Dump(lib)

	// lookup books by known author in the library
	// fmt.Println(lib.lookupByAuthor("Inline"))

	books := lib.lookupByAuthorName(jb.name)
	spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
	}

	// bookstest := lib.lookupByAuthorName("NilCase")
	// spew.Dump(bookstest)

	// // print out the first book's title and its author's name
	// if bookstest == nil {
	// 	fmt.Println("nil")
	// }

}
