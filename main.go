package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data dummy pengguna
var users = []map[string]interface{}{
	{"id": 1, "name": "Angga Jayadi"},
	{"id": 2, "name": "Jayadi Angga"},
}

// Handler adalah fungsi utama yang akan dipanggil oleh Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Inisialisasi router Gin
	router := gin.Default()

	// Endpoint GET /api/users
	router.GET("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})

	// Endpoint POST /api/users
	router.POST("/api/users", func(c *gin.Context) {
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

	// Jalankan router Gin sebagai handler
	router.ServeHTTP(w, r)
}
