func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		charSortedWord := sortWord(str)
		_, ok := mp[charSortedWord]
		if ok {
			mp[charSortedWord] = append(mp[charSortedWord], str)
		} else {
			mp[charSortedWord] = append([]string{str})
		}
	}
	var res [][]string
	for _, val := range mp {
		res = append(res, val)
	}

	return res
}

func sortWord(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func (i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
