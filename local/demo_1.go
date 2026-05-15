package main

import (
	"fmt"
	"unicode/utf8"
	"errors"
)

func main() {
	fmt.Println(utf8.RuneCountInString("Hello"))

	var numerator int = 11
	var denominator int = 3
	result, remainder, err  := intDivision(numerator, denominator)

	if (err != nil) {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result: %v\n", result)
	} else {
		fmt.Printf("Result: %v\nRemainder: %v\n", result, remainder)
	}

	var fixedSizeArray [3]int32 = [3]int32{1, 2, 3}
	anotherFixedSizeArray := [...]int32{4, 5, 6}
	
	fmt.Println(fixedSizeArray)
	fmt.Println(anotherFixedSizeArray)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Division by Zero!")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	
	return result, remainder,  err
}