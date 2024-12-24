package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFieldsSlice( returns int slice of fields
func GetFieldsSlice(str string) []int {
	strSlice := strings.Split(str, ",")
	intSlice := make([]int, 0, len(strSlice))
	for _, i := range strSlice {
		val, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, val)
	}
	return intSlice
}

func main() {
	f := flag.String("f", "", "\"fields\" - выбрать поля (колонки)")
	d := flag.String("d", "\t", "\"delimiter\" - использовать другой разделитель")
	s := flag.Bool("s", false, "\"separated\" - только строки с разделителем")
	flag.Parse()
	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		val := strings.Split(reader.Text(), *d)
		if len(val) > 1 {
			resStr := ""
			for _, idx := range GetFieldsSlice(*f) {
				if idx-1 < len(val) {
					resStr += val[idx-1] + *d
				}
			}
			fmt.Println(resStr)
			continue
		} else {
			if *s {
				continue
			} else {
				if len(val) > 0 {
					fmt.Println(val[0])
				}
			}
		}
	}
}
