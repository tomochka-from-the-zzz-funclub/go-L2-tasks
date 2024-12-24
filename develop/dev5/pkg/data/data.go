package filedata

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dev5/pkg/arguments"
)

type fileData struct {
	data []string
	arg  arguments.Arguments
}

func MakeFileData() *fileData {
	return &fileData{data: make([]string, 0)}
}

func (fd *fileData) FillFileData(filename string, arg arguments.Arguments) error {
	fd.arg = arg
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if *arg.Field_i {
			fd.data = append(fd.data, strings.ToLower(scanner.Text()))
		} else {
			fd.data = append(fd.data, scanner.Text())
		}
	}
	return nil
}

func (fd *fileData) Process() {
	if *fd.arg.Field_c {
		fd.Field_c()
	}
	if *fd.arg.Field_v {
		fd.Field_v()
	}
	if *fd.arg.Field_F {
		fd.Field_F()
	}
	if *fd.arg.Field_n {
		fd.Field_n()
	}
}

func (fd *fileData) Field_c() {
	fmt.Println("Flag c:")
	cnt := 0
	for _, str := range fd.data {
		if strings.Contains(str, fd.arg.Field_pattern) {
			cnt++
		}
	}
	if cnt == 0 {
		fmt.Println("There aren't strings containing this pattern")
	} else {
		fmt.Printf("Number of strings containing this pattern: %d\n", cnt)
	}
}

func (fd *fileData) Field_v() {
	fmt.Println("Flag v:")
	for _, str := range fd.data {
		if str != fd.arg.Field_pattern {
			fmt.Println(str)
		}
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (fd *fileData) Field_F() {
	fmt.Println("Flag F:")
	for i, str := range fd.data {
		if str == fd.arg.Field_pattern {
			fmt.Print("Match in string number ", i+1, ": \n")
			if *fd.arg.Field_C > 0 {
				for _, elem := range fd.data[Max(0, i-*fd.arg.Field_C):Min(len(fd.data), i+*fd.arg.Field_C+1)] {
					fmt.Println(elem)
				}
			} else if *fd.arg.Field_A > 0 {
				for _, elem := range fd.data[i:Min(len(fd.data), i+*fd.arg.Field_A+1)] {
					fmt.Println(elem)
				}
			} else if *fd.arg.Field_B > 0 {
				for _, elem := range fd.data[Max(0, i-*fd.arg.Field_C) : i+1] {
					fmt.Println(elem)
				}
			} else {
				fmt.Println(str)
			}
		}
	}
}

func (fd *fileData) Field_n() {
	fmt.Println("Flag n:")
	for i, str := range fd.data {
		if str == fd.arg.Field_pattern {
			fmt.Print("Match in string number: ", i, "\n")
		}
	}
}
