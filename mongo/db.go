package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://localhost:27017/gogam"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")
	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connect to", uri)
	Session = s
	Mongo = mongo
}