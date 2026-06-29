package validparentheses

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "()"
	expected := true
	result := isValid(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "()[]{}"
	expected := true
	result := isValid(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	s := "(]"
	expected := false
	result := isValid(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase4(t *testing.T) {
	s := "([])"
	expected := true
	result := isValid(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase5(t *testing.T) {
	s := "([)]"
	expected := false
	result := isValid(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
