package leetcode

func quicksort(arr *[]int, low, high int) {
	if low < high {
		i := partition(arr, low, high)
		quicksort(arr, low, i-1)
		quicksort(arr, i+1, high)
	}
}

func partition(arr *[]int, low, high int) int {
	prev := low - 1
	pivot := (*arr)[high-1]
	for current := 0; current < high; current++ {
		if (*arr)[current] <= pivot {
			prev++
			if prev < current {
				temp := (*arr)[current]
				(*arr)[current] = (*arr)[prev]
				(*arr)[prev] = temp
			}
		}
	}
	return prev
}