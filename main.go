package main

import (
	"github.com/auth0-community/go-auth0"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/godzyken/golang-angular/handlers"
	"gopkg.in/square/go-jose.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"time"
)

var (
	audience string
	domain   string
)

func main() {
	setAuth0Variables()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowOrigins:    []string{"*", "http://dev-c-559zpw.auth0.com"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "UPDATE"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
			"Accept-Encoding",
			"X-CSRF-Token",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	authorized := r.Group("/api")

	// Todos
	authorized.Use(authRequired())
	authorized.GET("/todo", handlers.GetTodoListHandler)
	authorized.OPTIONS("/todo", preflight)
	authorized.POST("/todo", handlers.AddTodoHandler)
	authorized.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	authorized.PUT("/todo", handlers.CompleteTodoHandler)

	// Songs
	authorized.POST("/song/createAsong", handlers.CreateAsong)

	go func() {
		err := r.Run("127.0.0.1:3000")
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceFul shutdown", sig)

}

func setAuth0Variables() {
	audience = os.Getenv("https://golang-angular-api/")
	domain = os.Getenv("dev-c-559zpw.auth0.com")
}

// ValidateRequest will verify that a token received from an http request
// is valid and signed by Auth0
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)
		_, err := validator.ValidateRequest(c.Request)
		if err != nil {
			log.Println(err)
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}
		c.Next()

	}
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
	return
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
