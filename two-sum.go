package main

func twoSum(nums []int, target int) []int {
	var res []int
	lookup := make(map[int]int)

	for i, num := range nums {
		index := target - num
		if _, ok := lookup[index]; ok {
			res = []int{lookup[index], i}
			return res
		}
		lookup[num] = i
	}
	return nil
}
