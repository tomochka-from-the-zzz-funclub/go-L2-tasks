package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-l2-tasks/develop/dev11/pkg/calendar"
	"github.com/go-l2-tasks/develop/dev11/pkg/event"
	"github.com/go-l2-tasks/develop/dev11/pkg/logger"
	"github.com/go-l2-tasks/develop/dev11/pkg/query"
)

var calendarData calendar.Calendar = *calendar.NewCalendar()

func StartServer(port string) error {
	// "/events_for_day?user_id=1&year=2010&month=10&day=10"
	mux := http.NewServeMux()
	mux.HandleFunc("/events_for_day", EventsForDay)
	mux.HandleFunc("/events_for_week", EventsForWeek)
	mux.HandleFunc("/events_for_month", EventsForMonth)

	mux.HandleFunc("/create_event", CreateEvent)
	mux.HandleFunc("/update_event", UpdateEvent)
	mux.HandleFunc("/delete_event", DeleteEvent)

	//err := http.ListenAndServe(port, logger)
	server := &http.Server{Addr: port, Handler: logger.RequestLogger(mux)}
	err := server.ListenAndServe()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		ctx := context.Background()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
	}()
	return err
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var NewEvent event.Event
	err := json.NewDecoder(r.Body).Decode(&NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	err = calendarData.AddEvent(NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
	}
	GoodResponseEvent(w, NewEvent, http.StatusOK)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var NewEvent event.Event
	err := json.NewDecoder(r.Body).Decode(&NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	err = calendarData.UpdateEvent(NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
	}
	GoodResponseEvent(w, NewEvent, http.StatusOK)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var NewEvent event.Event
	err := json.NewDecoder(r.Body).Decode(&NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	err = calendarData.DeleteEvent(NewEvent)
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
	}
	GoodResponseEvent(w, NewEvent, http.StatusOK)
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	q, err := EventsFor(w, r)
	if err != nil {
		return
	}
	var events []event.Event
	if events, err = calendarData.GetEventsForTheDay(*q); err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	GoodResponse(w, events, http.StatusOK)
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	q, err := EventsFor(w, r)
	if err != nil {
		return
	}
	var events []event.Event
	if events, err = calendarData.GetEventsForTheWeek(*q); err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	GoodResponse(w, events, http.StatusOK)
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	q, err := EventsFor(w, r)
	if err != nil {
		return
	}
	var events []event.Event
	if events, err = calendarData.GetEventsForTheMonth(*q); err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return
	}
	GoodResponse(w, events, http.StatusOK)
}

func EventsFor(w http.ResponseWriter, r *http.Request) (*query.Query, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return nil, errors.New("bad_response")
	}
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return nil, errors.New("bad_response")
	}
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return nil, errors.New("bad_response")
	}
	day, err := strconv.Atoi(r.URL.Query().Get("day"))
	if err != nil {
		BadResponse(w, err, http.StatusBadRequest)
		return nil, errors.New("bad_response")
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &query.Query{ID: userID, Date: date}, nil
}

type BadJson struct {
	Error string `error:"string"`
}

func BadResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	bj := BadJson{err.Error()}
	data, Error := json.Marshal(bj)
	if Error != nil {
		log.Fatal(Error)
	}
	w.Write(data)
}

type GoodJson struct {
	Result []event.Event `json:"result"`
}

func GoodResponse(w http.ResponseWriter, str []event.Event, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	bj := GoodJson{Result: str}
	data, Error := json.Marshal(bj)
	if Error != nil {
		log.Fatal(Error)
	}
	w.Write(data)
}

type GoodJsonEvent struct {
	Result event.Event `json:"result"`
}

func GoodResponseEvent(w http.ResponseWriter, str event.Event, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	bj := GoodJsonEvent{Result: str}
	data, Error := json.Marshal(bj)
	if Error != nil {
		log.Fatal(Error)
	}
	w.Write(data)
}
