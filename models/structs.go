package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

// Songs model
type Song struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}

type Songs []Song

// To do data structure for a task with a description of what to do
type Todo struct {
	ID       string `json:"_id" bson:"_id,omitempty"`
	Message  string `json:"message" bson:"message"`
	Complete bool   `json:"complete" bson:"complete"`
}

type Todos []Todo

type User struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" form:"username" binding:"required" bson:"username"`
	Password string        `json:"password" form:"password" binding:"required" bson:"password"`
	Email    string        `json:"email" form:"email" binding:"required" bson:"email"`
}

type Users []User

type Profile struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" form:"username" binding:"required" bson:"username"`
	Password string        `json:"password" form:"password" binding:"required" bson:"password"`
	Email    string        `json:"email" form:"email" binding:"required" bson:"email"`
}

func (s *Song) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*s).Title, (*s).Body)
}

func (t *Todo) String() string {
	return t.Message
}

func (u *User) String() string {
	return fmt.Sprintf("%d, de \"%s\"", (*u).ID, (*u).Username)
}

func (m *DbMongo) GetSong() (songs []Song, err error) {
	session := m.Session.Clone()
	defer session.Close()

	err = session.DB(m.SongsCollection).C("Song").Find(bson.M{}).All(&songs)
	return songs, err
}

func (m *DbMongo) PostSong(song *Song) (err error) {
	session := m.Session.Clone()
	defer session.Close()

	err = session.DB(m.SongsCollection).C("Song").Insert(&song)
	return err
}

func (m *DbMongo) GetTogo() (todos []Todo, err error) {
	session := m.Session.Clone()
	defer session.Close()

	err = session.DB(m.TodosCollection).C("Todo").Find(bson.M{}).All(&todos)
	return todos, err
}

func (m *DbMongo) PostTodo(todo *Todo) (err error) {
	session := m.Session.Clone()
	defer session.Close()

	err = session.DB(m.SongsCollection).C("Todo").Insert(&todo)
	return err
}
