package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMakeCommands(t *testing.T) {
	// Тест для команды pwd
	t.Run("pwd", func(t *testing.T) {
		var buf bytes.Buffer
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()

		originalStdout := os.Stdout // сохранить оригинальный stdout
		os.Stdout = w               // перенаправить stdout в pipe

		MakeCommands([]string{"pwd"})

		os.Stdout = originalStdout // восстановить stdout после теста
		io.Copy(&buf, r)           // скопировать данные из pipe в буфер

		currentDir, _ := os.Getwd()
		if buf.String() != currentDir+"\n" {
			t.Errorf("Expected %s but got %s", currentDir, buf.String())
		}
	})

	// Тест для команды echo
	t.Run("echo", func(t *testing.T) {
		var buf bytes.Buffer
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()

		originalStdout := os.Stdout
		os.Stdout = w

		MakeCommands([]string{"echo Hello, World!"})

		os.Stdout = originalStdout
		io.Copy(&buf, r)

		expected := "Hello, World!\n"
		if buf.String() != expected {
			t.Errorf("Expected %s but got %s", expected, buf.String())
		}
	})

	// Тест для команды cd
	t.Run("cd", func(t *testing.T) {
		originalDir, _ := os.Getwd() // сохранить оригинальный рабочий каталог
		defer os.Chdir(originalDir)  // восстановить рабочий каталог после теста

		testDir := "testdir"
		os.Mkdir(testDir, 0755)  // создать тестовую директорию
		defer os.Remove(testDir) // удалить тестовую директорию после теста

		MakeCommands([]string{"cd " + testDir}) // изменить на тестовую директорию

		currentDir, _ := os.Getwd()
		if currentDir != originalDir+"/"+testDir {
			t.Errorf("Expected to be in %s but got %s", originalDir+"/"+testDir, currentDir)
		}
	})

	// Тест для команды ps
	t.Run("ps", func(t *testing.T) {
		var buf bytes.Buffer
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()

		originalStdout := os.Stdout
		os.Stdout = w

		MakeCommands([]string{"ps"})

		os.Stdout = originalStdout
		io.Copy(&buf, r)

		if buf.String() == "" {
			t.Error("Expected output from ps command, but got none")
		}
	})

	// Тест для команды kill
	t.Run("kill", func(t *testing.T) {
		var buf bytes.Buffer
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()

		originalStdout := os.Stdout
		os.Stdout = w

		MakeCommands([]string{"kill 99999"}) // использование несуществующего PID

		os.Stdout = originalStdout
		io.Copy(&buf, r)

		if !strings.Contains(buf.String(), "no such process") {
			t.Errorf("Expected error message for non-existent PID, but got: %s", buf.String())
		}
	})
}
