// Link: https://leetcode.com/problems/left-and-right-sum-differences/?envType=daily-question&envId=2026-06-22
package leftandrightsumdifferences

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Complexity: O(n)
func leftRightDifference(nums []int) []int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	result := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			leftSum[i] = 0
		} else {
			leftSum[i] = nums[i-1] + leftSum[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightSum[i] = 0
		} else {
			rightSum[i] = nums[i+1] + rightSum[i+1]
		}
	}
	for i := range result {
		result[i] = abs(leftSum[i] - rightSum[i])
	}
	return result
}
