// Link: https://leetcode.com/problems/climbing-stairs
package climbingstarts

// Complexity: O(n), where n = len(memo)
func climbStairs(n int) int {
	memo := make([]int, n)
	return climbStairs2(n, memo)
}

func climbStairs2(n int, memo []int) int {
	if n <= 2 {
		return n
	}
	if memo[n-1] == 0 {
		memo[n-1] = climbStairs2(n-1, memo) + climbStairs2(n-2, memo)
	}
	return memo[n-1]
}
