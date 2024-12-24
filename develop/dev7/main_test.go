package main

import (
	"testing"
	"time"
)

// Тестируем функцию or
func TestOr(t *testing.T) {
	// Создаем сигналы с задержками
	sig1 := sig(100 * time.Millisecond)
	sig2 := sig(200 * time.Millisecond)
	sig3 := sig(300 * time.Millisecond)

	// Запускаем функцию or
	done := or(sig1, sig2, sig3)

	// Измеряем время до закрытия канала
	start := time.Now()
	<-done
	duration := time.Since(start)

	// Проверяем, что канал закрылся
	if duration < 100*time.Millisecond || duration > 150*time.Millisecond {
		t.Errorf("Expected done to close around 100-150ms, but it took %v", duration)
	}
}

// Функция для создания сигнала
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
