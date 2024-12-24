package arguments

import (
	"flag"
	"log"
)

type Arguments struct {
	Field_A       *int
	Field_B       *int
	Field_C       *int
	Field_c       *bool
	Field_i       *bool
	Field_v       *bool
	Field_F       *bool
	Field_n       *bool
	Field_pattern string
}

func MakeArguments() Arguments {
	a := Arguments{}
	a.Field_A = flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	a.Field_B = flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	a.Field_C = flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	a.Field_c = flag.Bool("c", false, "\"count\" (количество строк)")
	a.Field_i = flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	a.Field_v = flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	a.Field_F = flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	a.Field_n = flag.Bool("n", false, "\"line num\", напечатать номер строки")
	flag.Parse()
	var err error
	a.Field_pattern = flag.Arg(1)
	if err != nil {
		log.Fatal(err)
	}
	return a
}
