// Link: https://leetcode.com/problems/group-anagrams/
package groupanagrams

import (
	"maps"
	"slices"
)

// Complexity: O(m * n * log(n))
func groupAnagrams(strs []string) [][]string {
	strsMap := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		strsMap[sorted] = append(strsMap[sorted], s)
	}
	values := slices.Collect(maps.Values(strsMap))
	return values
}

func sortString(str string) string {
	chars := []rune(str)
	slices.Sort(chars)
	sortedString := string(chars)
	return sortedString
}
