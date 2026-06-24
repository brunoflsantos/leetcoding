package romantointeger

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "III"
	expected := 3
	result := romanToInt(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "LVIII"
	expected := 58
	result := romanToInt(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
func TestCase3(t *testing.T) {
	s := "MCMXCIV"
	expected := 1994
	result := romanToInt(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
