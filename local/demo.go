package main

import (
	"fmt"
	"math"
	"strconv"
)

const PERCENTAGE uint8 = 100

func main() {
	// Data Types
	// int => int8 - int64
	// float => float32 && float64
	// bool
	// byte => can hold character '$'
	// byte => int8
	// rune => int32
	// nil => undefined, null

	// Variables
	// Implicit var => var <variable_name> <data_type> = <value>
	// Explicit => <variable_name> := <value>
	// const

	// Helpful Go Packages
	// fmt -> Printf, Sprintf, Println
	// strconv (string to number converter) -> Atoi
	// math -> Math operations

	var number float64 = 250000
	percent := 23.25

	fmt.Printf("%.2f%% of %.2f = %.2f\n" , percent, number, CalculatePercentage(number, percent))

	unf := 12.5 / 0.23
	fmted := fmt.Sprintf("%.2f", unf)

	fmt.Printf("Unformatted: %f\n", unf)
	
	fmt.Printf("Formatted: %s\n", fmted)

	fmt.Println(math.Min(4, 5))

	stringToConvert := "1234"
	numberFromString, _ := strconv.Atoi(stringToConvert)

	fmt.Println(numberFromString)

	// Conditionals => if-else-if..., switch
	age := 50

	if age > 18 {
		fmt.Println("Approved")
	} else {
		fmt.Printf("Wait %d more years\n", 18 - age)
	}

	switch {
		case age == 10:	fmt.Println("A Decade")
		case age == 18:	fmt.Println(" Adult")
		case age == 20:	fmt.Println("Two Decades")
		case age == 25:	fmt.Println("Silver Jubilee")
		case age == 50:	fmt.Println("Golden Jubilee (Old)"); fallthrough
		case age == 100: fmt.Println("Century (Old)")
		case age == 1000, age % 1000 == 0: fmt.Println("Millenial (Old)")
	default: fmt.Println("Nothing Special about your age, dude!")
	}

	color := "purple"

	switch color {
		case "purple":
			fmt.Println("Correct! My Favorite color is Purple")
		case "yellow":
			fmt.Println("Yellow was ze former favorite color")
		default:
			fmt.Println("Wrong color")
	}

	// Loops => for,
	
}

func CalculatePercentage(number float64, percent float64) float64 {
	return (number * percent) / float64(PERCENTAGE)
}