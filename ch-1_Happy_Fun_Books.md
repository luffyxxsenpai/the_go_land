First this is how you print hello world in go
```go
package main
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

- every go program, needs a *main* package and a *main* function 
- so we have 2 things we primarily do in GO, we either create a library or an executable. package main tells GO to treat this code as a executable and not as a library, if the main is absent go will treat it as a library, only a "main" package can be used to create executables
- *func main* is called the entrypoint since GO compiler needs to know from where the program is starting, package main tells compiler to build it as an executable but func main tells the compiler to start the execution from this function and its fixed, we need a package main and a func main in our program.
- `go mod init books` mod files are simply the requirement files that contains the project name and most importantly all the externally used packages and their exact version used. it updates itself whenever we fetch any libs using `go get xyz`

```go
package main
import "fmt"
func main() {
    fmt.Println("Books in stock:")
    var title = "The City & The City"
    var author = "China Miéville"
    printBook(title, author)
    title = "The Sugar Barons"
    author = "Matthew Parker"
    printBook(title, author)
}
func printBook(title, author string) {
    fmt.Println(title, "by", author)
}
```
- in the above code, import "fmt" represents a built in go function used for formatting and manipulating strings
- to define a variable its `var var_name = var_value`
- go is smart, it can automatically infer the datatype of variables or we can explicitly define it 
- in the function printbook(title , author ) defines the input this function will have and the their type being string and its being called in our main function 

