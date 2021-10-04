package main

import (
	"fmt"
)

var taskList []Task

func main() {
	director := Director{"Main director"}
	id, err := director.giveTask("Do some stuff")
	id1, err1 := director.giveTask("Make a connection with database")
	fmt.Println(id, err)
	fmt.Println(id1, err1)
	fmt.Println(taskList)

	// fmt.Println(task.Deadline.Format("2006-01-02 15:04:05"))
}
