package main

import "fmt"

/*
 *  Chapter 2: Types
 */
func main() {
    //Number oparator
    fmt.Println("oparator")
    fmt.Println("1 + 1 =", 1 + 1)
    fmt.Println("1 - 1 =", 1 - 1)
    fmt.Println("1 / 1 =", 1 / 1)
    fmt.Println("2 * 2 =", 2 * 2)
    fmt.Println("\n")

    //Strings
    fmt.Println("Strings")
    fmt.Println("Len(): ", len("Hello, World"))
    fmt.Println("Position [1]: ", "Hello, World"[1])
    fmt.Println("Hello, " + "World")
    fmt.Println("\n")

    //Bool
    fmt.Println("Bool")
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println("\n")


    // Practice Exercises

    /* 1. Although overpowered for the task, you can use Go as a calculator. Write a program
     * that computes 32,132 × 42,452 and prints it to the terminal (use the * operator
     * for multiplication).
     */
    fmt.Println("Exercises")
    fmt.Println("32,132 * 42,452 = ", 32132 * 42452)
    fmt.Println("\n")
}


/** Theory Exercises
 *
 * 1. How are integers stored on a computer?
 *
 *    - On the computer the integers are stored on 2 digits (1 0) [Binary]
 *
 * 2. We know that (in base 10) the largest one-digit number is 9
 *    and the largest twodigit number is 99. Given that in binary the largest
 *    two-digit number is 11 (3), the largest three-digit number is 111 (7) and
 *    the largest four-digit number is 1111 (15), what’s the largest eight-digit number?
 *
 *    - 11111111 (255)
 *
 * 3. What is a string? How do you find its length?
 *
 *    - String is asequence of characters with a definite length used to represent text.
 *
 *    - You can find its length using a native function of Go "len".
 *      i.e.: len("Miguel")
 *
 *
 * 4. What’s the value of the expression (true && false) || (false && true) || !
 *   (false && false)?
 *
 *   - TRUE
 */
