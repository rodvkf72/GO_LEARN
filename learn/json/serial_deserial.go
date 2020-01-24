package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Deadline struct {
	time.Time
}

type NewDeadline(t time.Time) *Deadline{
  return &Deadline{t}
}

type status int

type Task struct {
	title    string
	status   status
	Deadline *Deadline
}

func Example_marshalJSON() {
	t := Task{
		"Laundry",
		2,
		NewDeadline(time.Date(2020, time.January, 24, 22, 46, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
