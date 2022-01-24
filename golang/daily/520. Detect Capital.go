package daily

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	const lowercaseIndex = 'a'

	var isLower bool

	if word[0] >= lowercaseIndex {
		isLower = true
	} else {
		isLower = word[1] >= lowercaseIndex
	}

	for i := 1; i < len(word); i++ {
		if isLower != (word[i] >= lowercaseIndex) {
			return false
		}
	}

	return true
}
