package tokopedia

func prison(n int32, m int32, h []int32, v []int32) int64 {
	// Write your code here
	// make slices with length of n and m
	horizontalRemoved := make([]bool, n)
	verticalRemoved := make([]bool, m)

	// On removed bar, put true on the slice's element
	for _, removeIndex := range h {
		horizontalRemoved[removeIndex-1] = true
	}
	for _, removeIndex := range v {
		verticalRemoved[removeIndex-1] = true
	}

	// Count max sequentially removed bar for horizontal
	var maxHorizontal, curHorizontal int32
	for i := 0; i < int(n); i++ {
		if horizontalRemoved[i] {
			curHorizontal++
			if curHorizontal > maxHorizontal {
				maxHorizontal = curHorizontal
			}
		} else {
			curHorizontal = 0
		}
	}

	// Count max sequentially removed bar for vertical
	var maxVertical, curVertical int32
	for i := 0; i < int(m); i++ {
		if verticalRemoved[i] {
			curVertical++
			if curVertical > maxVertical {
				maxVertical = curVertical
			}
		} else {
			curVertical = 0
		}
	}

	// Convert to long integer before multiplying
	horizontalGap := int64(maxHorizontal + 1)
	verticalGap := int64(maxVertical + 1)

	return horizontalGap * verticalGap
}
