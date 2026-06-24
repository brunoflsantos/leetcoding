package validpalindrome

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	result := isPalindrome(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "race a car"
	expected := false
	result := isPalindrome(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	s := " "
	expected := true
	result := isPalindrome(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
