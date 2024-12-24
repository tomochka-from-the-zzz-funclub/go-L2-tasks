package main

import (
	"errors"
	"testing"
	"time"
)

type MockNTPClient struct {
	mockTime time.Time
	mockErr  error
}

// Time возвращает замокированное время или ошибку
func (m *MockNTPClient) Time(host string) (time.Time, error) {
	return m.mockTime, m.mockErr
}

// Тесты
func TestNTPClient(t *testing.T) {
	tests := []struct {
		name     string
		mockTime time.Time
		mockErr  error
		expected time.Time
		hasErr   bool
	}{
		{
			name:     "successful time fetch",
			mockTime: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
			mockErr:  nil,
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
			hasErr:   false,
		},
		{
			name:     "error fetching time",
			mockTime: time.Time{},
			mockErr:  errors.New("failed to fetch time"),
			expected: time.Time{},
			hasErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockNTPClient{
				mockTime: tt.mockTime,
				mockErr:  tt.mockErr,
			}

			gotTime, err := client.Time("0.beevik-ntp.pool.ntp.org")

			if (err != nil) != tt.hasErr {
				t.Errorf("expected error: %v, got: %v", tt.hasErr, err)
			}
			if !gotTime.Equal(tt.expected) {
				t.Errorf("expected time: %v, got: %v", tt.expected, gotTime)
			}
		})
	}
}
