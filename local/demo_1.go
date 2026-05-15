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

	// Array Declarations
	var fixedSizeArray [3]int32 = [3]int32{1, 2, 3}
	anotherFixedSizeArray := [...]int32{4, 5, 6}
	
	fmt.Println(fixedSizeArray)
	fmt.Println(anotherFixedSizeArray)

	// Slices
	var intSlice []int32 = []int32{7, 8, 9}
	fmt.Println(intSlice)
	fmt.Printf("Length before Appending: %v\nCapacity: %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 10)
	fmt.Println(intSlice)
	fmt.Printf("Length after Appending: %v\nCapacity: %v\n", len(intSlice), cap(intSlice))

	// Slices 2
	var intSlice2 []int32 = make([]int32, 3, 8)	// `make` takes the slices' length and capacity (if omitted it just infers)
	fmt.Println(intSlice2)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap["Person 1"])	// default value when trying to access an invalid key in the map

	var myMap2 = map[string]uint8{"Person 1":23, "Person 2":30}
	fmt.Println(myMap2["Person 1"])

	var age, ok = myMap2["Person 2"]

	delete(myMap2, "Person 2")

	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Println("Data not found")
	}

	age, ok = myMap2["Person 2"]

	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Println("Data not found")
	}

	// Iteration
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