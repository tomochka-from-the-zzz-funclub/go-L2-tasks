package event

import (
	"encoding/json"
	"io"
	"time"
)

type Event struct {
	ID          int    `json:"id"`
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Day         int    `json:"day"`
	Hour        int    `json:"hour"`
	Minutes     int    `json:"minutes"`
	NameEvent   string `json:"nameevent"`
	Description string `json:"description"`
}

func (e *Event) Check() bool {
	if e.ID <= 0 {
		return false
	} else if e.Year <= 0 {
		return false
	} else if !(0 <= e.Month && e.Month <= 12) {
		return false
	} else if !(0 <= e.Day && e.Day <= 31) {
		return false
	} else if !(0 <= e.Hour && e.Hour <= 24) {
		return false
	} else if !(0 <= e.Minutes && e.Minutes <= 60) {
		return false
	} else if e.NameEvent == "" {
		return false
	}
	return true
}

func (e *Event) JsonDecode(r io.Reader) error {
	data := make([]byte, 0)
	r.Read(data)
	err := json.Unmarshal(data, &e)
	return err
}

func (e *Event) MakeDate() time.Time {
	return time.Date(e.Year, time.Month(e.Month), e.Day, e.Hour, e.Minutes, 0, 0, time.UTC)
}
