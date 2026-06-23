// Link: https://leetcode.com/problems/weighted-word-mapping
package weightedwordmapping

import "strings"

// Complexity: O(m * n), where m = len(words), n = len(words[i])
func mapWordWeights(words []string, weights []int) string {
	var result strings.Builder
	for _, s := range words {
		sum := 0
		for _, r := range s {
			i := int(r - 'a')
			weight := weights[i]
			sum += weight
		}
		result.WriteRune(rune('z' - sum%26))
	}
	return result.String()
}
