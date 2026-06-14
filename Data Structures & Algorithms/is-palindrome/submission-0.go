func isPalindrome(s string) bool {
	
	s = validStringProvider(s)
	left, right := 0, len(s) - 1
	fmt.Println(s)
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func validStringProvider(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res += strings.ToLower(string(str[i]))
		}
		if str[i] >= 'a' && str[i] <= 'z' {
			res += string(str[i])
		}
		if str[i] >= '0' && str[i] <= '9' {
			res += string(str[i])
		}
	}
	
	return res
}