package main

import "fmt"

//Control Structures
func main() {

	//For Loop even and odd example with if
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	var numberToLetter int32

	fmt.Println("Insert a number to covert it into a Word from 0 to 4")

	fmt.Scanf("%d", &numberToLetter)

	switch numberToLetter {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Not number found")
	}
}

/*
 * Exercise
 *
 *  What does the following program print?
 *      i := 10
 *      if i > 10 {
 *          fmt.Println("Big")
 *      } else {
 *          fmt.Println("Small")
 *      }
 *
 *  - Print the word 'Small' because the variable `i` is less of the number 10
 */
