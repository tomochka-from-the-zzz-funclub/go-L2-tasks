package arguments

import (
	"flag"
	"log"
	"os"
	"time"
)

type Arguments struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func MakeArguments() *Arguments {
	return &Arguments{}
}

func (a *Arguments) FillArguments() {
	timeoutRef := flag.String("timeout", "10s", "таймаут на подключение к серверу")
	flag.Parse()
	var err error
	a.Timeout, err = time.ParseDuration(*timeoutRef)
	if err != nil {
		log.Fatal(err)
	}
	a.Host = "127.0.0.1"
	if len(os.Args) > 1 {
		a.Host = flag.Arg(0)
		a.Port = flag.Arg(1)
	} else {
		a.Port = flag.Arg(0)
	}
}
