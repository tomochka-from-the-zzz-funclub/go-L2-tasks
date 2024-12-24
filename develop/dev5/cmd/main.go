package main

import (
	"flag"
	"log"

	"dev5/pkg/arguments"
	filedata "dev5/pkg/data"
)

func main() {
	arg := arguments.MakeArguments()
	fd := filedata.MakeFileData()
	err := fd.FillFileData(flag.Arg(0), arg)
	if err != nil {
		log.Fatal(err)
	}
	fd.Process()
}
