package filemap

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FileMap struct {
	data [][]string
}

func NewFileMap() *FileMap {
	return &FileMap{data: make([][]string, 0)}
}
func (fm *FileMap) FillFileMap(filename string) error {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fm.data = append(fm.data, strings.Fields(fileScanner.Text()))
	}
	return nil
}

func (fm *FileMap) deleteCopies() {
	newData := make([][]string, 0)
	setStrings := make(map[string]bool)
	for _, strArray := range fm.data {
		str := strings.Join(strArray, " ")
		if setStrings[str] {
			continue
		}
		setStrings[str] = true
		newData = append(newData, strings.Split(str, " "))
	}
	fm.data = newData
}

func (fm *FileMap) print(reverse bool) {
	if reverse {
		for i := len(fm.data) - 1; i >= 0; i-- {
			fmt.Println(fm.data[i])
		}
	} else {
		for _, str := range fm.data {
			fmt.Println(str)
		}
	}
}

func (fm *FileMap) Sort(k int, n, r, u bool) {
	if u {
		fm.deleteCopies()
	}
	k -= 1
	if n {
		sort.Slice(fm.data, func(i int, j int) bool {
			i_int, err := strconv.Atoi(fm.data[i][k])
			if err != nil {
				log.Fatal(err)
			}
			j_int, err := strconv.Atoi(fm.data[j][k])
			if err != nil {
				log.Fatal(err)
			}
			return i_int < j_int
		})
	} else {
		sort.Slice(fm.data, func(i int, j int) bool {
			return fm.data[i][k] < fm.data[j][k]
		})
	}
	if r {
		fm.print(true)
	} else {
		fm.print(false)
	}
}
