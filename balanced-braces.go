package main

func balanceBraces(s string) string {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--

			if count < 0 {
				s = s[:i] + s[i+1:]
				count++
			}
		}
	}

	if count == 0 {
		return s
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' && count > 0 {
			s = s[:i] + s[i+1:]
			//count--
		}
	}

	return s
}
