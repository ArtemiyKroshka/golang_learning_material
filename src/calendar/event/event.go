package event

import (
	"errors"
	"goCourse/src/calendar/date"
	"log"
)

type Event struct {
	title string
	date.Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(s string) {
	if len(s) > 15 {
		log.Fatal(errors.New("too long"))
	}
	e.title = s
}
