func isAnagram(s string, t string) bool {
	finalS := sortString(s)
	finalT := sortString(t)
	fmt.Println(finalS)
	fmt.Println(finalT)
	if finalS == finalT {
		return true
	}

	return false
}

func sortString (s string) string {
	chars := []rune(s)
	sort.Slice(chars, func (i, j int) bool {
		return chars[i]  < chars[j]
	})

	return string(chars)
}
