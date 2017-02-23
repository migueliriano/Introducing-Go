package main

import "fmt"

/*
 * Chapter 3: Variables
 */
func main() {
    //Inline Assignation
    var x string = "Hello, World"

    fmt.Println(x)

    var xDeclaration string
    xDeclaration = "String X declaration"

    fmt.Println(xDeclaration)

    // == is like + and return boolean
    var a string = "Hello"
    var b string = "World"

    fmt.Println(a == b) // false

    // shorter statement Declaration
    shorterVariable := "Hello, shorter statement"
    shorterVariableInt := 541
    fmt.Println(shorterVariable)
    fmt.Println(shorterVariableInt)

    //Multiple Variables
    var (
        var1 = 10
        var2 = 15
        var3 = 20
    )

    fmt.Println(var1)
    fmt.Println(var2)
    fmt.Println(var3)
}