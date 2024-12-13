package leetcode

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	output := ""
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			output += string(word1[i])
			i++
		}
		if j < len(word2) {
			output += string(word2[j])
			j++
		}
	}
	return output
}