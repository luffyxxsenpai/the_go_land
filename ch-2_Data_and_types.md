we have various types of data types like int, float, string, bool
- variable declaration can be done in various ways
	- `var book = "norwegian wood"` declares and define a value instantly
	- `var book string` defines a variable without any value 
	- data types have their zero values as their default value
		- string - "" empty quotations
		- int - 0
		- float - 0
		- bool - false 
	- when we assign a variable this is their default value unless we modify it 
	- type checks make sure we don't end up putting different kind of values in different data types like putting string in a int defined variable 

 structured data aka struct , we can group different kind of variable in a single variable around the same context like a book variable can have title, author, copies, price all combined in a single struct 
```go
type Book struct {
    Title  string
    Author string
    Copies int
}
```
now we can use it like 
```go
var book = Book {
	Title: "Never Let Me Go",
	Author: "Toshizuka Kawaguchi"
	Copies: 1
}

fmt.Println(book.Title, book.Author)
fmt.Printf("%v by %v - %v copies\n", book.Title, book.Author, book.Copies)
```

- fmt.println automatically put spaces while in fmt.printf we can control the output structure 
- heres the final code from ch2
```go
package main
import "fmt"

type Book struct {
	Title string
	Author string
	Copies int
}

var book = Book{
	Title: "Norwegian Wood",
	Author: "Haruki Murakami",
	Copies: 3,
}
func main() {
	PrintBook(book)
}
func PrintBook(mybooks Book) {
	fmt.Printf("%v by %v - copies %v\n", mybooks.Title, mybooks.Author, mybooks.Copies)
}
``` 