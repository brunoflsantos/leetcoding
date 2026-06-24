package leftandrightsumdifferences

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{10, 4, 8, 3}
	expected := []int{15, 1, 11, 22}
	result := leftRightDifference(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	expected := []int{0}
	result := leftRightDifference(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
