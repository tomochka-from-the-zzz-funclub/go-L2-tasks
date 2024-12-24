package main

import (
	"flag"

	filemap "dev3/pkg/fileMap"
)

func main() {
	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "обратная сортировка")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	fileData := filemap.NewFileMap()
	fileData.FillFileMap(flag.Arg(0))
	fileData.Sort(*k, *n, *r, *u)

}
