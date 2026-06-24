package containsduplicate

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := true
	result := containsDuplicate(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := false
	result := containsDuplicate(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected := true
	result := containsDuplicate(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
