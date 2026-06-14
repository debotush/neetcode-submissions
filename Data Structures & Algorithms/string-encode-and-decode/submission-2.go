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

func getSeparatorProvider (str string, startingPosition int) (string, int) {
	separator := ""
	for i := startingPosition + 1; i > 0 ; i++ {
		if str[i] == '|' {
			break
		}
		separator += string(str[i])
	}  

	return separator, len(separator) // l:15
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
		if encoded[i] == '|' {
			separator, separatorLength := getSeparatorProvider(encoded, i) 
			stringLength :=  getStringLength(separator)
			stringStartingPosition := i + separatorLength + 1
			j := stringStartingPosition
			var str string  
			for j <= stringStartingPosition + stringLength {
				if string(encoded[j]) != "|" && stringStartingPosition != j {
					str += string(encoded[j])
				}

				j++
			}
			decodedStringList = append(decodedStringList, str)
			i = j 
		}
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