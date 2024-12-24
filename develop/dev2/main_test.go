package main

import (
	"testing"
)

func TestUnpackageOfString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"45", "", true},                     // Ошибка: строка не должна начинаться с числа без символа
		{"a4bc2d5e", "aaaabccddddde", false}, // Пример с буквами и цифрами
		{"abcd", "abcd", false},              // Без повторений
		{"", "", false},                      // Пустая строка
		{"qwe\\4\\5", "qwe45", false},        // Экранирование
		{"qwe\\\\5", "qwe\\\\\\\\\\", false}, // Два обратных слеша
		{"qwe\\45", "qwe44444", false},       // Экранирование с цифрой
	}

	for _, test := range tests {
		result, err := UnpackageOfString(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %q, got result %q", test.input, result)
			}
		} else {
			if err != nil {
				t.Errorf("UnpackageOfString(%q) returned an error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("UnpackageOfString(%q) = %q; expected %q", test.input, result, test.expected)
			}
		}
	}
}
