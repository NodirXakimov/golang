package main

import "errors"

type TeamLeader struct {
	name string
}

func (tl TeamLeader) deligateTask(id int, taskStep string) (err error) {
	for _, value := range taskList {
		if value.ID == id {
			if value.Status != "done" {
				err := errors.New("Task already has been done")
				return err
			}
			value.Status = taskStep
			break
		}
	}
	return
}
