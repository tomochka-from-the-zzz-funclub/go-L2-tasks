package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// TimeFetcher интерфейс для получения времени
type TimeFetcher interface {
	Time(host string) (time.Time, error)
}

// NTPClient реализует TimeFetcher
type NTPClient struct{}

// Time получает текущее время от NTP-сервера
func (n *NTPClient) Time(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	client := &NTPClient{}
	time, err := client.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := fmt.Fprint(os.Stderr, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
