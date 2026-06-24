package palindromenumber

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	x := 121
	expected := true
	result := isPalindrome(x)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	x := -121
	expected := false
	result := isPalindrome(x)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	x := 10
	expected := false
	result := isPalindrome(x)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
