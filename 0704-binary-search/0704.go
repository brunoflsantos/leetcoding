// Link: https://leetcode.com/problems/binary-search/
package binarysearch

// Complexity: O(log n)
func search(nums []int, target int) int {
	return _search(nums, 0, len(nums)-1, target)
}

func _search(nums []int, i int, j int, target int) int {
	if i > j { // Element not found
		return -1
	}
	midIndex := i + (j-i)/2
	midElement := nums[midIndex]
	if target == midElement {
		return midIndex
	} else if target > midElement {
		return _search(nums, midIndex+1, j, target)
	} else {
		return _search(nums, i, midIndex-1, target)
	}
}
