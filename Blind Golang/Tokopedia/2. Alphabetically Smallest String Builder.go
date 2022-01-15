package tokopedia

import (
	"strings"
)

func smallestString(substrings []string) string {
	// Write your code here
	for i := 0; i < len(substrings); i++ {
		for j := i + 1; j < len(substrings); j++ {
			// Sort array with custom comparison
			if substrings[i]+substrings[j] > substrings[j]+substrings[i] {
				substrings[i], substrings[j] = substrings[j], substrings[i]
			}
		}
	}

	return strings.Join(substrings, "")
}
