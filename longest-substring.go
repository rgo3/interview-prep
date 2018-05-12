package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) (res int) {
	n := len(s)
	index := 0

	var chars [128]int

	for i := 0; i < n; i++ {
		index = int(math.Max(float64(chars[s[i]]), float64(index)))
		res = int(math.Max(float64(res), float64(i-index+1)))
		chars[s[i]] = i + 1
	}
	return
}

func lengthOfLongestSubstring2(s string) (res int) {
	i := 0
	j := 0
	chars := make(map[byte]bool)
	for i < len(s) && j < len(s) {
		if ok := chars[s[j]]; !ok {
			chars[s[j]] = true
			j++
			res = int(math.Max(float64(res), float64(j-i)))
		} else {
			chars[s[j]] = false
			i++
		}
	}
	return
}
