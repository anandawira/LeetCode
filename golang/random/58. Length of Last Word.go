package random

func lengthOfLastWord(s string) int {
	endIndex := len(s)-1

	for endIndex >= 0 {
		if s[endIndex] == ' ' {
			endIndex--
		} else {
			break
		}
	}

	startIndex := endIndex - 1

	for startIndex >= 0 {
		if s[startIndex] != ' ' {
			startIndex--
		} else {
			break
		}
	}

	return endIndex - startIndex
}
