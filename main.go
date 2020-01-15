package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping Successful",
		})
	})
	r.GET("/device-definitions/:id", readDefinition)

	r.POST("/devices-definitions/create", createDefinition)

	r.PUT("/device-definitions/:id", updateDefinition)

	r.DELETE("/device-definitions/:id", deleteDefinition)

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":8004"))
}
