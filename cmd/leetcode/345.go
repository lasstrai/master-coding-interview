package leetcode

func isVowel(c byte) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true,
		'o': true, 'u': true,
		'A': true, 'E': true, 'I': true,
		'O': true, 'U': true,
	}
	return vowels[c]
}

func reverseVowels(s string) string {
	word := []byte(s)
	left := 0
	right := len(word) - 1
	for left < right {
		for !isVowel(word[left]) {
			left++
		}
		for !isVowel(word[right]) {
			right--
		}
		temp := word[left]
		word[left] = word[right]
		word[right] = temp
		left++
		right--
	}
	return string(word)
}
