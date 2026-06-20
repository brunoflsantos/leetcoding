// Link: https://leetcode.com/problems/two-sum/
package twosum

// Complexity: O(n)
func twoSum(nums []int, target int) []int {
	numsMapped := make(map[int]int)
	for i, num := range nums {
		numToSearch := target - num
		if k, ok := numsMapped[numToSearch]; ok {
			return []int{k, i}
		}
		numsMapped[num] = i
	}
	return nil
}
