package countingbits

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// ans[i/2] takes the previous result already calculated
		// (i % 2) adds 1 if the number is odd
		ans[i] = ans[i/2] + (i % 2)
	}
	return ans
}
