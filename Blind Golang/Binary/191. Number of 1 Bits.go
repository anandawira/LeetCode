package binary

// Shifting mask
func hammingWeight(num uint32) int {
	mask := 1
	var result int
	for i := 0; i < 32; i++ {
		if num&uint32(mask<<i) > 0 {
			result++
		}
	}
	return result
}

// Shifting the number
// func hammingWeight(num uint32) int {
// 	var result int
// 	for num > 0 {
// 		if num&1 == 1 {
// 			result++
// 		}
// 		num = num >> 1
// 	}

// 	return result
// }
