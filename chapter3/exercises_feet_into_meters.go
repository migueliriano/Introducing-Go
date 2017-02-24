package main

import "fmt"

/*
 * Write another program that converts from feet into meters (1 ft = 0.3048 m).
 */
func main() {

	var feet float32

	const feetToMeter float32 = 0.3048

	fmt.Println("\n");

	fmt.Println("---------converts from feet into meters---------\n")

	fmt.Println("Enter feet: \n")

	fmt.Scanf("%f", &feet)

	meters := feet * feetToMeter

	fmt.Println("\n");

	fmt.Println("In meters is:", meters);
}