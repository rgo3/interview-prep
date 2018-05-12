package main

func reverse(s string) string {
	c := []rune(string(s))

	for i := 0; i < len(s)/2; i++ {
		c[i], c[len(s)-i-1] = c[len(s)-i-1], c[i]
	}

	return string(c)
}
