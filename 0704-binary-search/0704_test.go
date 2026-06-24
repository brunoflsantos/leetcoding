package binarysearch

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	result := search(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	expected := -1
	result := search(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
