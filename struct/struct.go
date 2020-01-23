package main

import "time"

type Task struct {
	title string
	done  bool
	due   *time.Time
}

var myTask = Task{"laundry", false, nil}

/*var myTask = Task{
  title: "laundry",
  done: true, <--important rest
}*/

type status int

type Task2 struct {
	title2 string
	status status
	due2   *time.Time
}

const (
	UNKNOWN status = iota
	TODO
	DONE
	/*
		UNKNOWN status = 0
		TODO    status = 1
		DONE    status = 2
	*/
)
