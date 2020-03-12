package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/godzyken/golang-angular/models"
	"github.com/godzyken/golang-angular/todo"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	_ "time"
)

func MiddleDB(m *models.DbMongo) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := m.SetSession()
		if err != nil {
			c.Abort()
		} else {
			c.Set("mongo", m)
			c.Next()
		}
	}
}

// To Do
func GetTodoListHandler(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*models.DbMongo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't reach db", "body": nil})
	}
	c.JSON(http.StatusOK, todo.Get())

	t, err := mongo.GetTogo()

	// fmt.Printf("\ntodo: %v, ok: %v\n", todo, ok)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't get a todo from database", "body": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "get a todo success", "body": t})
	}
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

// Song post
func CreateAsong(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*models.DbMongo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't connect to db", "body": nil})
	}
	var req models.Song

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect data", "body": nil})
		return
	} else {
		err := mongo.PostSong(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error post to db", "body": nil})
		}
		c.JSON(http.StatusOK, gin.H{"message": "post a song success", "body": req})
	}
}

// List all songs
func GetAsong(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*models.DbMongo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't reach db", "body": nil})
	}

	song, err := mongo.GetSong()
	// fmt.Printf("\nsong: %v, ok: %v\n", song, ok)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't get a song from database", "body": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "get a song success", "body": song})
	}
}

// Update an article
func Update(c *gin.Context) {

}

// Delete an article
func Delete(c *gin.Context) {

}

//--------------------------END Song-------------------//
