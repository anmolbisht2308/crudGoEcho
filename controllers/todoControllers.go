package controllers

import (
	"context"
	"net/http"
	"time"

	"myapp/config"
	"myapp/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTodo handles the creation of a new todo
func CreateTodo(c echo.Context) error {
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todo.ID = primitive.NewObjectID()

	collection := config.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, todo)
}

// GetTodos retrieves all todos
func GetTodos(c echo.Context) error {
	var todos []models.Todo

	collection := config.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}

// GetTodo retrieves a single todo by ID
func GetTodo(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	collection := config.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var todo models.Todo
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Todo not found")
	}

	return c.JSON(http.StatusOK, todo)
}

// UpdateTodo updates a todo by ID
func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := config.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": todo})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	collection := config.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
