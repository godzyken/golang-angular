package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionSongs holds the name of the articles collection
	CollectionSong = "songs"
	//CollectionUser = "users"
)

// Songs model
type Song struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}

type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" form:"username" binding:"required" bson:"username"`
	Password string        `json:"password" form:"password" binding:"required" bson:"password"`
	Email    string        `json:"email" form:"email" binding:"required" bson:"email"`
}

type Profile struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" form:"username" binding:"required" bson:"username"`
	Password string        `json:"password" form:"password" binding:"required" bson:"password"`
	Email    string        `json:"email" form:"email" binding:"required" bson:"email"`
}
