package main

import (
	"reflect"
	"testing"
)

func TestMakeAnagramMap(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{"листок": []string{"слиток", "столик"}, "пятак": []string{"пятка", "тяпка"}},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			input:    []string{"пятак", "пятка", "тяпка", "пятак", "пятак", "пятак", "пятак"},
			expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}},
		},
		{
			input:    []string{"пятАк", "Пятка", "тяпКа", "листок", "слиток", "столик"},
			expected: map[string][]string{"листок": []string{"слиток", "столик"}, "пятак": []string{"пятка", "тяпка"}},
		},
	}

	for _, test := range tests {
		actual := MakeAnagramMap(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, actual)
		}
	}
}
