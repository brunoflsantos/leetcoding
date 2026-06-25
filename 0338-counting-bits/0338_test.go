package countingbits

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	n := 2
	expected := []int{0, 1, 1}
	result := countBits(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	n := 5
	expected := []int{0, 1, 1, 2, 1, 2}
	result := countBits(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
