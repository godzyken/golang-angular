package router

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
	//validator *auth0.JWTValidator
	//AdminGroup = "Admin"
)

func main() {
	setAuth0Variables()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	//r.Use(CORSMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowOrigins:    []string{"*", "http://dev-c-559zpw.auth0.com"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "UPDATE"},
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

	//r.OPTIONS("/*path", CORSMiddleware())
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

	authorized := r.Group("/")

	// Todos
	authorized.Use(authRequired())
	authorized.GET("/todo", handlers.GetTodoListHandler)
	authorized.POST("/todo", handlers.AddTodoHandler)
	authorized.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	authorized.PUT("/todo", handlers.CompleteTodoHandler)

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

	// User Functions - located in users.go
	//user := r.Group("/user")
	//user.Use(AuthRequired())
	//{
	//	user.GET("/info", userInfo)
	//	user.PUT("/update", updateUser)
	//	user.DELETE("/delete", deleteUser)
	//}

	//manager := manage.NewDefaultManager()
	//
	////use mongoDB token store
	//manager.MapTokenStorage(
	//	mongo.NewTokenStore(mongo.NewConfig(
	//		"mongodb://127.0.0.1:27017",
	//		"oauth2",
	//	)),
	//)

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://dev-c-559zpw.auth0.com")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.JSON(http.StatusOK, struct{}{})

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		} else {
			c.Next()
		}
	}
}
