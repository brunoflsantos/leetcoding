package pascaltriangle

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	numRows := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	result := generate(numRows)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	numRows := 1
	expected := [][]int{
		{1},
	}
	result := generate(numRows)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
