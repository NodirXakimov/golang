package main

import (
	"fmt"
)

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

var tasks []Task

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Do you check your progress")

	// Init Router
	// r := mux.NewRouter()

	// Route Handlers / Endpoints
	// r.HandleFunc("api/")
}
