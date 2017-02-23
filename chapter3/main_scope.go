package main

import "fmt"

/*
 * Chapter 3: Scope
 */
var x string = "Hello, World" //Other functions can access to this variable

func main() {
    //Inline Assignation
    fmt.Println(x)
}

func printx() {
    fmt.Println(x)
}