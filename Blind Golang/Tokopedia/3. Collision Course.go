package tokopedia

func collision(speed []int32, pos int32) int32 {
	// Write your code here
	// Get out main particle's speed
	mainSpeed := speed[pos]

	var result int32
	// Count particle on left with higher speed
	for i := 0; i < int(pos); i++ {
		if speed[i] > mainSpeed {
			result++
		}
	}
	// Count particle on right with lower speed
	for i := int(pos + 1); i < len(speed); i++ {
		if speed[i] < mainSpeed {
			result++
		}
	}

	return result
}
