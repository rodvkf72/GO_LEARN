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

func Example_marshalJSON() {
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

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Bilk","Status":2,"Deadline":"2020-01-25T15:40:00Z"}`) //not ''
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
}

func main() {
	Example_marshalJSON()
	Example_unmarshalJSON()
}

//*important : JSON package serializes only uppercase fields
