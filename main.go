package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data dummy pengguna
var users = []map[string]interface{}{
	{"id": 1, "name": "Angga Jayadi"},
	{"id": 2, "name": "Jayadi Angga"},
}

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Endpoint GET /api/users
	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})

	// Endpoint POST /api/users
	r.POST("/api/users", func(c *gin.Context) {
		var newUser map[string]interface{}

		// Bind JSON body ke variabel newUser
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Tambahkan pengguna baru ke slice users
		users = append(users, newUser)

		// Kembalikan respons dengan data pengguna yang baru ditambahkan
		c.JSON(http.StatusCreated, gin.H{
			"data": newUser,
		})
	})

	// Jalankan server di port 8080
	r.Run()
}
