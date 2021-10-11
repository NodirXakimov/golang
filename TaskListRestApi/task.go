package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Priority  string `json:"priority"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	DueDate   string `json:"due_date"`
}

var tasks = []Task{
	{"1", "First task", "init", "high", "10-10-2021 15:15:15", "1", "1"},
	{"2", "Second task", "dev", "middle", "10-10-2021 15:15:15", "2", "2"},
}

// Get all tasks
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// Get single task
func getTask(c *gin.Context) {
	id := c.Param("id")
	// Loop through tasks and find by id
	for _, task := range tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Create new task
func createTask(c *gin.Context) {
	var task Task

	// ID generated - not safe
	task.ID = strconv.Itoa(len(tasks) + 1)

	// Attached created time to the task
	t := time.Now()
	task.CreatedAt = t.String()

	// Attached default 3 day to the DueDate ( .AddDate(year, month, day) )
	t = time.Now().AddDate(0, 0, 3)
	task.DueDate = t.String()

	// Call BindJSON to bind the received JSON to task.
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	// Add the new task to the slice.
	tasks = append(tasks, task)

	c.IndentedJSON(http.StatusCreated, task)
}

// Update a task
func updateTask(c *gin.Context) {
	id := c.Param("id")
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			if err := c.BindJSON(&task); err != nil {
				c.IndentedJSON(http.StatusBadRequest, err)
				return
			}
			task.ID = id // Mock ID - not safe
			tasks = append(tasks, task)
			c.IndentedJSON(http.StatusAccepted, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, tasks)
}

// Delete a task
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	isDelete := false
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			isDelete = true
			break
		}
	}
	if isDelete {
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}

func main() {

	router := gin.Default()
	router.GET("/api/tasks", getTasks)
	router.POST("/api/tasks", createTask)
	router.GET("/api/tasks/:id", getTask)
	router.PUT("/api/tasks/:id", updateTask)
	router.DELETE("/api/tasks/:id", deleteTask)

	log.Fatal(router.Run(":8080"))
}
