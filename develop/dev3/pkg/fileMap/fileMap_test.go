package filemap

import (
	"os"
	"testing"
)

func TestNewFileMap(t *testing.T) {
	fm := NewFileMap()
	if fm == nil {
		t.Fatal("Expected a new FileMap instance, got nil")
	}
	if len(fm.data) != 0 {
		t.Fatalf("Expected empty data slice, got %v", fm.data)
	}
}

func TestFillFileMap(t *testing.T) {
	// Создаем временный файл для тестирования
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // Удаляем файл после теста

	// Записываем тестовые данные
	testData := "line1\nline2\nline1\nline3\n"
	if _, err := tmpFile.WriteString(testData); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close() // Закрываем файл, чтобы его можно было открыть для чтения

	fm := NewFileMap()
	err = fm.FillFileMap(tmpFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedData := [][]string{
		{"line1"},
		{"line2"},
		{"line1"},
		{"line3"},
	}

	if len(fm.data) != len(expectedData) {
		t.Fatalf("Expected data length %d, got %d", len(expectedData), len(fm.data))
	}

	for i, line := range fm.data {
		if len(line) != len(expectedData[i]) {
			t.Fatalf("Expected line length %d, got %d", len(expectedData[i]), len(line))
		}
	}
}

func TestDeleteCopies(t *testing.T) {
	fm := NewFileMap()
	fm.data = [][]string{
		{"line1"},
		{"line2"},
		{"line1"},
		{"line3"},
	}

	fm.deleteCopies()

	expectedData := [][]string{
		{"line1"},
		{"line2"},
		{"line3"},
	}

	if len(fm.data) != len(expectedData) {
		t.Fatalf("Expected data length %d, got %d", len(expectedData), len(fm.data))
	}

	for i, line := range fm.data {
		if len(line) != len(expectedData[i]) {
			t.Fatalf("Expected line length %d, got %d", len(expectedData[i]), len(line))
		}
	}
}

func TestSort(t *testing.T) {
	fm := NewFileMap()
	fm.data = [][]string{
		{"2"},
		{"3"},
		{"1"},
	}

	fm.Sort(1, false, false, false)

	expectedData := [][]string{
		{"1"},
		{"2"},
		{"3"},
	}

	if len(fm.data) != len(expectedData) {
		t.Fatalf("Expected data length %d, got %d", len(expectedData), len(fm.data))
	}

	for i, line := range fm.data {
		if line[0] != expectedData[i][0] {
			t.Fatalf("Expected line %v, got %v", expectedData[i], line)
		}
	}
}

func TestPrint(t *testing.T) {
	// Для тестирования функции print можно использовать вывод в буфер
	// или просто проверить, что функция не вызывает паники.
	fm := NewFileMap()
	fm.data = [][]string{
		{"line1"},
		{"line2"},
	}

	// Здесь просто проверяем, что функция не вызывает паники
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Print panicked: %v", r)
		}
	}()
	fm.print(false)
}
