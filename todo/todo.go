package todo

import (
	"errors"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   Todo               `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

// To do data structure for a task with a description of what to do
type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

//type Jwks struct {
//	Keys []jose.JSONWebKey `json:"keys"`
//}
//
//type JSONWebKeys struct {
//	Kty string `json:"kty"`
//	Kid string `json:"kid"`
//	Use string `json:"use"`
//	N   string `json:"n"`
//	E   string `json:"e"`
//	W5c string `json:"x5c"`
//}

var (
	list []Todo
	mtx  sync.RWMutex
)

func init() {
	list = []Todo{}
}

func Get() []Todo {
	return list
}

func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

// Delete will remove a To do from the Todo list
func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

// Complete will set the complete boolean to true, marking a todo as
// completed
func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	return nil
}

func newTodo(msg string) Todo {
	return Todo{
		ID:       xid.New().String(),
		Message:  msg,
		Complete: false,
	}
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find todo based on id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
