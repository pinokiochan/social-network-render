package utils

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"example@example.com", true},          // Валидный email
		{"example@example", false},             // Нет TLD
		{"example@.com", false},                // Нет имени хоста
		{"example@example..com", false},        // Два символа ".." в домене
		{"example@ex..ample.com", false},       // Два символа ".." в домене
		{"example@ex-ample.com", true},         // Валидный email с дефисом в домене
		{"ex@ample.com", true},                 // Валидный email
		{"example@example.com.br", true},       // Валидный email с длинным TLD
		{"example@verylongtld.com", false},     // TLD слишком длинный
		{"", false},                           // Пустой email
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			result := IsValidEmail(test.email)
			if result != test.expected {
				t.Errorf("For email %s, expected %v, but got %v", test.email, test.expected, result)
			}
		})
	}
}
