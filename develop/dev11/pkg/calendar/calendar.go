package calendar

import (
	"errors"
	"sync"
	"time"

	"github.com/go-l2-tasks/develop/dev11/pkg/event"
	"github.com/go-l2-tasks/develop/dev11/pkg/query"
)

type Calendar struct {
	m    sync.RWMutex
	data map[int]map[time.Time]event.Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		m:    sync.RWMutex{},
		data: make(map[int]map[time.Time]event.Event),
	}
}

func (c *Calendar) AddEvent(ev event.Event) error {
	c.m.Lock()
	defer c.m.Unlock()
	_, ok := c.data[ev.ID]
	if !ok {
		c.data[ev.ID] = make(map[time.Time]event.Event)
	}
	date := ev.MakeDate()
	_, ok = c.data[ev.ID][date]
	if ok {
		return errors.New("the event in this day already exists, ERROR 400")
	}
	c.data[ev.ID][date] = ev
	return nil
}

func (c *Calendar) UpdateEvent(newEvent event.Event) error {
	c.m.Lock()
	defer c.m.Unlock()
	_, ok := c.data[newEvent.ID]
	if !ok {
		return errors.New("there aren't events for this, ERROR 400")
	}
	date := newEvent.MakeDate()
	_, ok = c.data[newEvent.ID][date]
	if !ok {
		return errors.New("there aren't events in this day, ERROR 400")
	}
	c.data[newEvent.ID][date] = newEvent
	return nil
}

func (c *Calendar) DeleteEvent(event event.Event) error {
	c.m.Lock()
	defer c.m.Unlock()
	_, ok := c.data[event.ID]
	if !ok {
		return errors.New("there aren't events for this, ERROR 400")
	}
	date := event.MakeDate()
	_, ok = c.data[event.ID][date]
	if !ok {
		return errors.New("there aren't events in this day, ERROR 400")
	}
	delete(c.data[event.ID], date)
	return nil
}

func (c *Calendar) GetEventsForTheDay(q query.Query) ([]event.Event, error) {
	return c.GetEvents(q, 24*time.Hour)
}

func (c *Calendar) GetEventsForTheWeek(q query.Query) ([]event.Event, error) {
	return c.GetEvents(q, 7*24*time.Hour)
}

func (c *Calendar) GetEventsForTheMonth(q query.Query) ([]event.Event, error) {
	return c.GetEvents(q, 30*24*time.Hour)
}

func (c *Calendar) GetEvents(q query.Query, t time.Duration) ([]event.Event, error) {
	answ := make([]event.Event, 0)
	c.m.RLock()
	_, ok := c.data[q.ID]
	if !ok {
		return answ, errors.New("there aren't events for this id, ERROR 400")
	}
	defer c.m.RUnlock()
	for _, rec := range c.data[q.ID] {
		if rec.MakeDate().Sub(q.Date) <= t {
			answ = append(answ, rec)
		}
	}
	return answ, nil
}
