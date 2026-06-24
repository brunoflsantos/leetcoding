package integertoroman

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	num := 3749
	expected := "MMMDCCXLIX"
	result := intToRoman(num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	num := 58
	expected := "LVIII"
	result := intToRoman(num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	num := 1994
	expected := "MCMXCIV"
	result := intToRoman(num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
