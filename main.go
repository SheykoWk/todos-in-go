package main

import (
	"github.com/SheykoWk/todos-in-go.git/controllers"
	"github.com/SheykoWk/todos-in-go.git/formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	formatter.NewJSONFormatter()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "ok",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "ok",
		})
	})
	r.GET("/api/todos", controllers.GetTodos)
	r.POST("/api/todos", controllers.CreateTodo)
	r.DELETE("/api/todos/:id", controllers.DeleteTodo)
	r.PUT("/api/todos/:id", controllers.UpdateTodo)
	r.Run()
}
