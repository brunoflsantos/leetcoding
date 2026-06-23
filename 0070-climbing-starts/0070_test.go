package climbingstarts

import (
	"reflect"
	"testing"
)

func TestGenerate1(t *testing.T) {
	n := 2

	expected := 2

	result := climbStairs(n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestGenerate2(t *testing.T) {
	n := 3

	expected := 3

	result := climbStairs(n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestGenerate3(t *testing.T) {
	n := 5

	expected := 8

	result := climbStairs(n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
