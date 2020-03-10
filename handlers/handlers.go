package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/godzyken/golang-angular/todo"
	"golang-angular/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// To Do
func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todo.Get())
}

// Addtodo will add a message on the list
func AddTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	//err = c.BindJSON(&json.Encoder{})
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": todo.Add(todoItem.Message)})
}

// DeleteTodoHandler will delete a specified to do based on user http input
func DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")
	if err := todo.Delete(todoID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// CompleteTodoHandler will complete a specified to do based on user http input
func CompleteTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if todo.Complete(todoItem.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todo.Todo, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return todo.Todo{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToTodo(body)
}

func convertJSONBodyToTodo(jsonBody []byte) (todo.Todo, int, error) {
	var todoItem todo.Todo
	err := json.Unmarshal(jsonBody, &todoItem)
	if err != nil {
		return todo.Todo{}, http.StatusBadRequest, err
	}
	return todoItem, http.StatusOK, nil
}

//-------------------END TO DO---------------------------//

// Song
func CreateAsong(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	song := models.Song{}
	err := c.Bind(&song)
	if err != nil {
		c.Error(err)
		return
	}

	song.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	song.UpdatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionSong).Insert(song)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all songs
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	songs := []models.Song{}
	err := db.C(models.CollectionSong).Find(nil).Sort("-_id").All(&songs)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, songs)
}

// Update an article
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	song := models.Song{}
	err := c.Bind(&song)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": song.Id}
	doc := bson.M{
		"title":      song.Title,
		"body":       song.Body,
		"created_on": song.CreatedOn,
		"updated_on": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionSong).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an article
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	song := models.Song{}
	err := c.Bind(&song)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": song.Id}
	err = db.C(models.CollectionSong).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}

//--------------------------END Song-------------------//
