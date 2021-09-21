package main

import "fmt"

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

var contacts []Contact
var tasks []Task

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Do you check your progress")
}
