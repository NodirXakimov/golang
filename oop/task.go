package main

import "time"

type Task struct {
	ID        int
	Name      string
	Status    string
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
