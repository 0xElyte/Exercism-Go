package isbnverifier

// IsValidISBN determines if a string is a valid ISBN-10.
func IsValidISBN(isbn string) bool {
	sum := 0
	count := 10

	for i := 0; i < len(isbn); i++ {
		c := isbn[i]

		if c == '-' {
			continue
		}

		if c >= '0' && c <= '9' {
			sum += int(c-'0') * count
			count--
		} else if c == 'X' && count == 1 {
			sum += 10
			count--
		} else {
			return false
		}
	}

	return count == 0 && sum%11 == 0
}