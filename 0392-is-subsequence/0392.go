// Link: https://leetcode.com/problems/is-subsequence
package issubsequence

// Complexity: O(n), where n = len(t)
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	result := false
	for i < len(s) {
		result = false
		for j < len(t) {
			if s[i] == t[j] {
				i++
				j++
				result = true
				break
			}
			j++
		}
		if result == false {
			return false
		}
	}
	return true
}
