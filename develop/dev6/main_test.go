package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Флаги
var (
	f = flag.String("f", "", "\"fields\" - выбрать поля (колонки)")
	d = flag.String("d", "\t", "\"delimiter\" - использовать другой разделитель")
	s = flag.Bool("s", false, "\"separated\" - только строки с разделителем")
)

// Функция для выполнения основного кода с подменой аргументов
func runWithArgs(args []string, input string) (string, error) {
	var outBuffer bytes.Buffer
	// Устанавливаем значение для флагов
	os.Args = args
	flag.Parse()

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), *d)
		if len(val) > 1 {
			resStr := ""
			for _, idx := range GetFieldsSlice(*f) {
				if idx-1 < len(val) {
					resStr += val[idx-1] + *d
				}
			}
			// Удалить последний разделитель, если он есть
			if len(resStr) > 0 {
				resStr = resStr[:len(resStr)-len(*d)]
			}
			fmt.Fprintln(&outBuffer, resStr)
		} else {
			if *s {
				continue
			} else {
				if len(val) > 0 {
					fmt.Fprintln(&outBuffer, val[0])
				}
			}
		}
	}
	return outBuffer.String(), scanner.Err()
}

func TestGetFieldsSlice(t *testing.T) {
	tests := []struct {
		input  string
		output []int
	}{
		{"1,2,3", []int{1, 2, 3}},
		{"10,20,30", []int{10, 20, 30}},
		{"-1,-2,-3", []int{-1, -2, -3}},
	}

	for _, test := range tests {
		got := GetFieldsSlice(test.input)
		if !equal(got, test.output) {
			t.Errorf("GetFieldsSlice(%q) = %v; want %v", test.input, got, test.output)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMain(t *testing.T) {
	tests := []struct {
		input     string
		fields    string
		delimiter string
		separated bool
		expected  string
	}{
		{
			input:     "1,2,3\n4,5,6\n7,8,9\n",
			fields:    "1,3",
			delimiter: ",",
			separated: false,
			expected:  "1,3\n4,6\n7,9\n",
		},
		{
			input:     "apple\tbanana\tcherry\nberry\tmelon\tpeach\n",
			fields:    "1,2",
			delimiter: "\t",
			separated: false,
			expected:  "apple\tbanana\nberry\tmelon\n",
		},
		{
			input:     "one\ttwo\nthree\tfour\nfive\tsix\n",
			fields:    "1",
			delimiter: "\t",
			separated: true,
			expected:  "one\nthree\nfive\n",
		},
		{
			input:     "no\tnot\tsep\nyes\t\tseparated\n\tyes\n",
			fields:    "1",
			delimiter: "\t",
			separated: true,
			expected:  "no\nyes\n\n",
		},
	}

	for _, test := range tests {
		args := []string{"cmd", "-f", test.fields, "-d", test.delimiter}
		if test.separated {
			args = append(args, "-s")
		}
		got, err := runWithArgs(args, test.input)
		if err != nil {
			t.Errorf("Error running test: %v", err)
		}
		if got != test.expected {
			t.Errorf("runWithArgs(%q) = %q; want %q", test.input, got, test.expected)
		}
	}
}
