package daily

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	start, end := 0, len(s1)-1

	target := make(map[byte]int)
	current := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		target[s1[i]]++
	}

	for i := start; i <= end; i++ {
		if _, exist := target[s2[i]]; exist {
			current[s2[i]]++
		}
	}

	if isSameMap(target, current) {
		return true
	}

	for end < len(s2)-1 {
		if _, exist := target[s2[start]]; exist {
			current[s2[start]]--
		}
		start++

		end++
		if _, exist := target[s2[end]]; exist {
			current[s2[end]]++
		}

		if isSameMap(target, current) {
			return true
		}
	}
	return false
}

func isSameMap(target, current map[byte]int) bool {
	isEqual := true
	for key, val := range target {
		if current[key] != val {
			isEqual = false
			break
		}
	}

	return isEqual
}
