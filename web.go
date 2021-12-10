package main

import (
	"strconv"

	"dh.com/test1/app"
	"github.com/gin-gonic/gin"
)

type AddTodoParams struct {
	Description string `uri:"description" binding:"required"`
}

type CompleteTodoParams struct {
	Id string `uri:"id" binding:"required"`
}

func main() {
	db := &app.InMemDB{}
	service := app.NewTodoService(db)

	r := gin.Default()

	r.GET("/list", func(c *gin.Context) {
		todos := service.List()
		c.JSON(200, gin.H{
			"todos": todos,
		})
	})

	r.GET("/add/:description", func(c *gin.Context) {
		var params AddTodoParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		todo := service.Add(params.Description)
		c.JSON(200, gin.H{"todo": todo})
	})

	r.GET("/complete/:id", func(c *gin.Context) {
		var params CompleteTodoParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		i, err := strconv.Atoi(params.Id)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		todo := service.Complete(i)
		c.JSON(200, gin.H{
			"todo": todo,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
