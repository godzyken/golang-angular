package models

type Reader interface {
	AllSongs() (Songs, error)
	SongFromID(ID int) (Song, error)
	AllTodos() (Todos, error)
	UserId(ID int) (User, error)
	AllUsers() (Users, error)
}
