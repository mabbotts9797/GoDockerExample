package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

var (
	toDoList []*toDo = make([]*toDo, 0)
	flake            = sonyflake.NewSonyflake(sonyflake.Settings{})
)

func main() {
	r := gin.Default()

	//GET todos
	r.GET("/todo", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"toDos": toDoList,
		})
	})

	//POST add a todo
	r.POST("/todo", func(c *gin.Context) {
		var toDo toDo

		if c.ShouldBind(&toDo) == nil {

			//Generate a uid
			id, _ := flake.NextID()

			toDo.ID = id
			toDoList = append(toDoList, &toDo)
			c.Status(200)
			return
		}

		c.JSON(400, "Failed to bind to model")
	})

	//DELETE delete a todo
	r.DELETE("/todo/:id", func(c *gin.Context) {

		//Loop through in memory list and remove matching id if exists.
		if id := c.Param("id"); id != "" {

			idUint, err := strconv.ParseUint(id, 10, 64)

			//Parsing errors are caught here, exit loop if failed to parse id
			if err != nil {
				c.JSON(400, "Failed to parse ID")
				return
			}

			for in, val := range toDoList {

				if val.ID == idUint {

					toDoList = append(toDoList[:in], toDoList[in+1:]...)
				}
			}
		}

		c.Status(200)
	})

	//PUT update a todo
	r.PUT("/todo/:id", func(c *gin.Context) {

		//Loop through in memory list and update matching id if exists.
		if id := c.Param("id"); id != "" {

			idUint, err := strconv.ParseUint(id, 10, 64)

			var toDo toDo

			//Parsing and model binding errors are caught here, exit loop if failed to parse id
			if c.ShouldBind(&toDo) != nil || err != nil {
				c.JSON(400, "Failed to parse ID")
				return
			}

			for in, val := range toDoList {

				if val.ID == idUint {
					toDo.ID = idUint
					toDoList[in] = &toDo
				}
			}
		}

		c.Status(200)
	})

	r.Run()
}

//ToDo type
type toDo struct {
	ID          uint64
	Title       string `form:"title"`
	Description string `form:"description"`
}
