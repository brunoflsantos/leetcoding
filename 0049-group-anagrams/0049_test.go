package groupanagrams

import (
	"reflect"
	"slices"
	"testing"
)

// normalizeGroups normalize the result such that the test accepts results with any order
func normalizeGroups(groups [][]string) [][]string {
	result := make([][]string, len(groups))
	for i, g := range groups {
		sorted := slices.Clone(g)
		slices.Sort(sorted)
		result[i] = sorted
	}
	slices.SortFunc(result, func(a, b []string) int {
		return slices.Compare(a, b)
	})
	return result
}

func TestCase1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(normalizeGroups(result), normalizeGroups(expected)) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	strs := []string{""}
	expected := [][]string{{""}}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(normalizeGroups(result), normalizeGroups(expected)) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	strs := []string{"a"}
	expected := [][]string{{"a"}}
	result := groupAnagrams(strs)
	if !reflect.DeepEqual(normalizeGroups(result), normalizeGroups(expected)) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
