package leetcode

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	a := len(str1)
	b := len(str2)
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return str1[0:a]
}