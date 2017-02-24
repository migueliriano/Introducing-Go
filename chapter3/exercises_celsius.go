package main

import "fmt"

/*
 * Using the example program as a starting point, write a program that converts
 * from Fahrenheit into Celsius (C = (F − 32) * 5/9).
 */
func main() {

    var fahrenheit int64

    fmt.Println("\n");

    fmt.Println("---------from Fahrenheit into Celsius---------\n")

    fmt.Println("Enter a Fahrenheit: \n")

    fmt.Scanf("%d", &fahrenheit)

    celsius := (fahrenheit - 32) * 5 / 9

    fmt.Println("\n");

    fmt.Println("In Celsius is:", celsius);
}