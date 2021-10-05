package main

import (
	"time"
)

type Director struct {
	name string
}

func (d *Director) giveTask(title string) (id int, er error) {
	task := Task{
		ID:        len(taskList) + 1,
		Name:      title,
		Status:    "initial",
		Deadline:  time.Now().Add(259200),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	taskList = append(taskList, task)
	return task.ID, er
}
