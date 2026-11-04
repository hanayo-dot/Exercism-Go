package luhn

func Valid(id string) bool {
	// Remove spaces and validate characters.
	digits := make([]int, 0, len(id))

	for _, r := range id {
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		digits = append(digits, int(r-'0'))
	}

	if len(digits) <= 1 {
		return false
	}

	sum := 0
	double := false

	// Work from right to left, doubling every second digit.
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}