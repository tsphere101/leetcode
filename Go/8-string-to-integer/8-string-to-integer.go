package leetcode

func myAtoi(s string) int {
	// digit chars 48 to 57 : c > 47 c < 58
	// alphabet 65-90 and 97-122

	result := 0
	negative := false
	hasPositiveSign := false
	end := false
	for _, c := range s {
		// white space
		if string(c) == " " && end {
			break
		}
		// char is digit
		if c >= 48 && c <= 57 {
			result = result*10 + (int(c) - 48)
			if !negative && result > 2147483647 {
				result = 2147483647
				break
			}
			if negative && result > 2147483648 {
				result = 2147483648
				break
			}

			end = true
			continue
		}

		if end {
			break
		}

		// if any -
		if string(c) == "-" {
			if negative {
				return 0
			}
			negative = true
		}
		if string(c) == "+" {
			hasPositiveSign = true
			end = true
		}

		// if c is alphabet
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			result = 0
			break
		}
		// if c is .
		if string(c) == "." {
			result = 0
			break
		}

	}
	if negative && hasPositiveSign {
		result = 0
	}
	if negative {
		result = -result
	}
	return result
}
