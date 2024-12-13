package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	output := make([]bool, len(candies))
	for j, value := range candies {
		if value+extraCandies >= max {
			output[j] = true
		}
	}
	return output
}