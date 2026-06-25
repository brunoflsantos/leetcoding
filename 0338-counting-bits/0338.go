package countingbits

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := range result {
		result[i] = countSetBits(i)
	}
	return result
}

// countSetBits implements the Brian Kernighan's Algorithm for counting the 1's in
// the bit representation of a number
func countSetBits(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
