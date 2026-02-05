package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

var book = Book{
	Title:  "Norwegian Wood",
	Author: "Haruki Murakami",
	Copies: 3,
}

func main() {
	PrintBook(book)
}

func PrintBook(mybooks Book) {
	fmt.Printf("%v by %v - copies %v\n", mybooks.Title, mybooks.Author, mybooks.Copies)
}