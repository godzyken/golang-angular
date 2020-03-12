package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/mgo.v2"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	rnd_str := make([]rune, n)
	for i := range rnd_str {
		rnd_str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rnd_str)
}

type DbMongo struct {
	Host            string
	Port            string
	Addrs           string
	Server          string
	DbName          string
	SongsCollection string
	TodosCollection string
	UsersCollection string
	Info            *mgo.DialInfo
	Session         *mgo.Session
}

func (m *DbMongo) SetDefault() {
	m.Host = "localhost"
	m.Addrs = "localhost:27017"
	m.DbName = "context"
	//m.EventTTLAfterEnd = 1 * time.Second
	//m.StdEventTTL = 20 * time.Minute
	m.Info = &mgo.DialInfo{
		Addrs:    []string{m.Addrs},
		Timeout:  60 * time.Second,
		Database: m.DbName,
	}
}

func (m *DbMongo) Drop() (err error) {
	session := m.Session.Clone()
	defer session.Close()

	err = session.DB(m.DbName).DropDatabase()
	if err != nil {
		return err
	}
	return nil
}

func (m *DbMongo) Init() (err error) {
	err = m.Drop()
	if err != nil {
		fmt.Printf("\n drop database error: %v\n", err)
	}

	todo := Todo{}
	todo.Message = rndStr(8)
	err = m.PostTodo(&todo)

	return err
}

func (m *DbMongo) SetSession() (err error) {
	m.Session, err = mgo.DialWithInfo(m.Info)
	if err != nil {
		m.Session, err = mgo.Dial(m.Host)
		if err != nil {
			return err
		}
	}
	return err
}

// Get
func (m *DbMongo) AllSongs() (Songs, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.SongsCollection)
	var result []Song
	if err = c.Find(nil).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *DbMongo) SongFromID(ID int) (Song, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return Song{}, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.SongsCollection)
	var result Song
	if err = c.Find(bson.M{"id": ID}).One(&result); err != nil {
		return Song{}, err
	}
	return result, nil
}

func (m *DbMongo) AllTodos() (Todos, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.TodosCollection)
	var result []Todo
	if err = c.Find(nil).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *DbMongo) AllUsers() (Users, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.UsersCollection)
	var result []User
	if err = c.Find(nil).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *DbMongo) UserFromID(ID int) (User, error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return User{}, err
	}
	defer session.Close()

	c := session.DB(m.DbName).C(m.UsersCollection)
	var result User
	if err = c.Find(bson.M{"id": ID}).One(&result); err != nil {
		return User{}, err
	}
	return result, nil
}

// Post
// SaveUser
func (m *DbMongo) SaveUserID(u User) error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(m.DbName).C(m.UsersCollection)
	return c.Insert(u)
}

//SaveSong prent un Song et l'enregistre en base
func (m *DbMongo) SaveSong(s Song) error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(m.DbName).C(m.SongsCollection)
	return c.Insert(s)
}

//SaveTodo prent un To do et l'enregistre en base
func (m *DbMongo) SaveTodo(t Todo) error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(m.DbName).C(m.TodosCollection)
	return c.Insert(t)
}
