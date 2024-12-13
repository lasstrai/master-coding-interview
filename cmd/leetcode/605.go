package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, value := range flowerbed {
		if value == 0 {
			cond1 := i == 0 || flowerbed[i-1] == 0
			cond2 := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if cond1 && cond2 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}