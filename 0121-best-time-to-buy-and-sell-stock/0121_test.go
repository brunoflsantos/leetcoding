package besttimetobuyandsellstock

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	result := maxProfit(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0
	result := maxProfit(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
