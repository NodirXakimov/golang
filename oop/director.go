package main

import (
	"math/rand"
	"time"
)

type Director struct {
	name string
}

func (d *Director) giveTask(title string) (id int, er error) {
	task := Task{
		ID:        rand.Intn(1000),
		Name:      title,
		Status:    "initial",
		Deadline:  time.Now().Add(259200),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	taskList = append(taskList, task)
	return task.ID, er
}
