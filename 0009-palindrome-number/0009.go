// Link: https://leetcode.com/problems/palindrome-number/
package palindromenumber

import "strconv"

// Complexity: O(n)
func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	for i, j := 0, len(xStr)-1; i < j; i, j = i+1, j-1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}
