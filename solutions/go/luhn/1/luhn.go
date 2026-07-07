package luhn


func Valid(id string) bool {
	sum := 0
	count := 0

	
	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]

		
		if c == ' ' {
			continue
		}

		
		if c < '0' || c > '9' {
			return false
		}

		count++
		val := int(c - '0')

		
		if count%2 == 0 {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}

		sum += val
	}

	
	return count > 1 && sum%10 == 0
}