package main

import (
	"errors"
)

type Developer struct {
	name string
}

func (d *Developer) develop(id int, taskStep string) (err error) {
	for index, value := range taskList {
		if value.ID == id {
			if value.Status != "initial" {
				err := errors.New("Developer can only develop initial statused task")
				return err
			}
			taskList[index].Status = taskStep
			break
		}
	}
	return err
}
