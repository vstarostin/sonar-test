package main

import (
	"log"
	"os"

	"github.tools.sap/project-piper/azure-demo-k8s-go/pkg/types"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	
	users := types.DefaultUsers()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})
	
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:<port> (for windows "localhost:<port>")
}
