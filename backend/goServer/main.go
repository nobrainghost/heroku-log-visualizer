package main

import (
	"net/http"

	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type LogEntry struct {
	ID        int       `json:"id"`
	Source    string    `json:"source"`
	TimeStamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
