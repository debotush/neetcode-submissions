type Solution struct{}

// functions I need, 
// separatorProvider(stringNumber, stringLength)
// separatorRemovalFunc()
// At first, I need a function that will provide me the separator
//
// Then, I need to concatanate all the strings using separator, provided by separatorProvider
// Then implement Encode.
// At last step, I need to implement the Decode.

func separatorProvider(stringLength int) string {
	return fmt.Sprintf("|l:%d|", stringLength)
}

func getSeparatorProvider(str string, startingPosition int) (string, int) {
    separator := ""
    for i := startingPosition + 1; i < len(str); i++ {  // ← i < len(str), NOT i > 0
        if str[i] == '|' {
            fullSepLen := i - startingPosition + 1  // includes both '|' chars
            return separator, fullSepLen
        }
        separator += string(str[i])
    }
    return separator, len(separator)
}

func (s *Solution) Encode(strs []string) string {
	encodedString := ""
	for i := 0; i < len(strs); i++ {
		encodedString += separatorProvider(len(strs[i]))
		encodedString += strs[i]
	}

	return encodedString
}


// dam, babua
// | l : 3 | d a m | l : 5 | b a b u a
// 0 1 2 3 4 5 6 7 8 91011121314151617

func (s *Solution) Decode(encoded string) []string {
	i := 0
	decodedStringList := make([]string, 0)
	for i < len(encoded) {
		separator, fullSepLen := getSeparatorProvider(encoded, i)
		stringLength := getStringLength(separator)
		start := i + fullSepLen
		decodedStringList = append(decodedStringList, encoded[start:start+stringLength])
		i = start + stringLength
	}
	return decodedStringList
}


func getStringLength(str string) int {
	resStr := ""
	flag := false
	for i := 0 ; i < len(str); i++ {
		if flag {
			resStr += string(str[i])
		}
		if str[i] == ':' {
			flag = true
		}
	}
	res, _ := strconv.Atoi(resStr)
	return res
}