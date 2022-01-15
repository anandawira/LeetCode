package tokopedia

func slowestKey(keyTimes [][]int32) string {
	// Write your code here
	keyDuration := make(map[int32]int32)

	var prevTime int32
	for _, pair := range keyTimes {
		// Extract key and time from the pair
		key := pair[0]
		time := pair[1]

		curDuration := time - prevTime
		prevDuration := keyDuration[key]

		if curDuration > prevDuration {
			keyDuration[key] = curDuration
		}

		prevTime = time
	}

	var slowestKey int32
	var slowestDuration int32

	// Get slowest key from the map
	for key, duration := range keyDuration {
		if duration > slowestDuration {
			slowestKey = key
			slowestDuration = duration
		}
	}

	return string([]byte{byte(slowestKey + 97)})
}
