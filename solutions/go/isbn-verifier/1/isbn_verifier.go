package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	// Remove hyphens.
	isbn = strings.ReplaceAll(isbn, "-", "")

	// ISBN-10 must contain exactly 10 characters.
	if len(isbn) != 10 {
		return false
	}

	sum := 0

	for i, r := range isbn {
		var value int

		if i == 9 && r == 'X' {
			value = 10
		} else if unicode.IsDigit(r) {
			value = int(r - '0')
		} else {
			return false
		}

		sum += value * (10 - i)
	}

	return sum%11 == 0
}

