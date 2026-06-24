package findthehighestaltitude

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	gain := []int{-5, 1, 5, 0, -7}
	expected := 1
	result := largestAltitude(gain)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	expected := 0
	result := largestAltitude(gain)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
