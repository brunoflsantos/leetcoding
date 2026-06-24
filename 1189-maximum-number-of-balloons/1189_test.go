package maximumnumberofballoons

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	text := "nlaebolko"
	expected := 1
	result := maxNumberOfBalloons(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	text := "loonbalxballpoon"
	expected := 2
	result := maxNumberOfBalloons(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	text := "leetcode"
	expected := 0
	result := maxNumberOfBalloons(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
