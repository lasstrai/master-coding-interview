package leetcode

// 1. Two Sum

func twoSums(nums []int, target int) []int {
	seens := make(map[int]int)
	for currentIndex, value := range nums {
		prevIndex, exists := seens[value]
		if exists {
			return []int{prevIndex, currentIndex}
		} else {
			diff := target - value
			seens[diff] = currentIndex
		}
	}
	return []int{0, 0}
}