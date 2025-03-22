package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Angga"},
	{ID: 2, Name: "Jayadi"},
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/users", getUsers)
	r.POST("/api/users", createUser)
	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	route := setupRouter()
	route.ServeHTTP(w, r)
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	r.Run(":" + port)
}
