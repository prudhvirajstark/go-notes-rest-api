package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port ", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("Failed to start server : ", err)
	}
}
