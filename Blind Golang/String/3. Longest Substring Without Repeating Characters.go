package string

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	max := 1
	curSub := s[:0]
	for i := 1; i < len(s); i++ {
		isExist := false
		for j := 0; j < len(curSub); j++ {
			if s[i] == curSub[j] {
				curSub = curSub[j+1:]
				curSub += string(s[i])
				isExist = true
				break
			}
		}

		if !isExist {
			curSub += string(s[i])
			if len(curSub) > max {
				max = len(curSub)
			}
		}
	}

	return max
}
