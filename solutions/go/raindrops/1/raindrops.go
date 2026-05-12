package raindrops

import "fmt"

func Convert(number int) string {
	result := ""
	found := false

	if (number % 3) == 0 {
		result += "Pling"
		found = true
	}

	if (number % 5) == 0 {
		result += "Plang"
		found = true
	}

	if (number % 7) == 0 {
		result += "Plong"
		found = true
	}

	if !found {
		result = fmt.Sprintf("%d", number)
	}

	return result
}