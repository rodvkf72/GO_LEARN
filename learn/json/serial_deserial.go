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

type status int

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func main() {
	t := Task{
		Title:    "Laundry",
		Status:   2,
		Deadline: NewDeadline(time.Date(2020, time.January, 24, 22, 46, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
