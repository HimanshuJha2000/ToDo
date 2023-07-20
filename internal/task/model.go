package task

import (
	"errors"
	"time"
)

// Task represents a TODO task
type Task struct {
	ID        int
	Title     string
	Content   string
	Deadline  time.Time
	Tags      []string
	Completed bool
}

// ValidateTask validates the received task
func (task *Task) ValidateTask() error {
	if task.Title == "" {
		return errors.New("task title cannot be empty")
	}
	if task.Deadline.String() == "" {
		return errors.New("task deadline cannot be empty")
	}
	if task.Deadline.Before(time.Now()) {
		return errors.New("deadline cannot be in past date")
	}
	return nil
}
