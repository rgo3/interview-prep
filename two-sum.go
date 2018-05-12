package main

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, num := range nums {
		index := target - num
		if _, ok := lookup[index]; ok {
			res := [2]int{lookup[index], i}
			return res[:2]
		}
		lookup[num] = i
	}
	return make([]int, 2)
}
