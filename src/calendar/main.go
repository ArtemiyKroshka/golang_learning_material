package main

import (
	"fmt"
	"goCourse/src/calendar/event"
)

func main() {
	date := event.Event{}
	date.SetTitle("love marina")
	date.SetDay(20)
	date.SetMonth(4)
	date.SetYear(2023)
	fmt.Println(date)

}
