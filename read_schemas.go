package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readDefinition(c echo.Context) error {
	var schema DeviceSchema
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	err = DB.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&schema)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, schema)
}

func readAllDefinitions(c echo.Context) error {
	var schemas []DeviceSchema
	userID := c.Param("id")
	filter := bson.D{{"uid", userID}}
	cur, err := DB.Collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var schema DeviceSchema
		err := cur.Decode(&schema)
		if err != nil {
			panic(err)
		}
		schemas = append(schemas, schema)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return c.JSON(200, schemas)
}