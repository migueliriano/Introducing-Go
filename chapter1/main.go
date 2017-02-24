package main

import "fmt"

/*
 * Chapter 1: Getting Started
 */
func main() {
    fmt.Println("Hello, World")

    /*  Practice Exercises
     *
     *  1. Modify the program we wrote so that instead of printing Hello, World it prints
     *  Hello, my name is followed by your name.
     */
    fmt.Println("Hello, my name is: Miguel Angel Liriano")
}

/** Theory Exercises
 *
 * 1. What is whitespace?
 *
 *    - It's a character does not correspond to a visible mark, but typically
 *      does occupy an area on a page.
 *
 * 2. What is a comment? What are the two ways of writing a comment?
 *
 *    - The comment is the way you can add the description and documentation
 *      of what your programn do The commands are not execute.
 *
 *    - We can write inline and block comments.
 *
 * 3. Our program began with package main. What would the files in the fmt package
 *    begin with?
 *
 *    - package fmt
 *
 * 4. We used the Println function defined in the fmt package. If you wanted to use
 *    the Exit function from the os package, what would you need to do?
 *
 *    - We need to defined os in the begin of this file e.i: import "os"
 */