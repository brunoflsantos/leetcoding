package pascaltriangle2

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	rowIndex := 3
	expected := []int{1, 3, 3, 1}
	result := getRow(rowIndex)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	rowIndex := 0
	expected := []int{1}
	result := getRow(rowIndex)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	rowIndex := 1
	expected := []int{1, 1}
	result := getRow(rowIndex)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
