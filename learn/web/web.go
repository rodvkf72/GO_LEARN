package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/rodvkf72/learn/stru"
)

type ID string

//DataAccess is an interface to access tasks
type DataAccess interface {
	Get(id ID) (stru.Task, error)
	Put(id ID, t stru.Task) error
	Post(t stru.Task) (ID, error)
	Delete(id ID) error
}

//MemoryDataAccess is a simple in-memory database
type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID int64
}

//NewMemoryDataAccess returns a new NewMemoryDataAccess
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}
}

//ErrTaskNotExist occurs when the task with the ID was not found
var ErrTaskNotExist = errors.New("task does not exist")

//Get returns a task with a given ID
func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exist {
		return task.Task{}, ErrTaskNotExist
	}
	return t, nil
}

//Put updates a task with a given ID with t
func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

//Post adds a new task
func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

//Delete removes the task with a given ID
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

//ResponseError is the error for the JSON Response
type ResponseError struct {
	Err error
}

//MarshalJSON returns the JSON representation of the error
func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

//UnmarshallJSON parses the JSON representation of the error
func (err *ResponseError) UnmarshalJSON(b []byte) err {
	var v interface{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNotExist.Error() {
			err.Err = ErrTaskNotExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

type Response struct {
	ID    ID            `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}

var m = NEwMemoryDataAccess()

const pathPrefix = "/api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World2 !")
}
func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}
