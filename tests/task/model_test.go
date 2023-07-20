package task_test

import (
	"github.com/MyOrg/ToDo/internal/task"
	"testing"
	"time"
)

func TestValidateTask(t *testing.T) {
	// valid task
	t.Run("Valid Task", func(t *testing.T) {
		validTask := &task.Task{
			ID:        1,
			Title:     "Sample Task",
			Content:   "Task content",
			Deadline:  time.Now().Add(time.Hour), // Set deadline to be one hour in the future
			Tags:      []string{"tag1", "tag2"},
			Completed: false,
		}

		err := validTask.ValidateTask()
		if err != nil {
			t.Errorf("Expected no error, but got: %s", err.Error())
		}
	})

	// task with an empty title
	t.Run("Task with Empty Title", func(t *testing.T) {
		emptyTitleTask := &task.Task{
			ID:        2,
			Title:     "", // Empty title
			Content:   "Task content",
			Deadline:  time.Now().Add(time.Hour),
			Tags:      []string{"tag1", "tag2"},
			Completed: false,
		}

		err := emptyTitleTask.ValidateTask()
		expectedErr := "task title cannot be empty"
		if err == nil {
			t.Error("Expected an error, but got no error")
		} else if err.Error() != expectedErr {
			t.Errorf("Expected error message: %s, but got: %s", expectedErr, err.Error())
		}
	})

	// task with a past deadline
	t.Run("Task with Past Deadline", func(t *testing.T) {
		pastDeadlineTask := &task.Task{
			ID:        3,
			Title:     "Sample Task",
			Content:   "Task content",
			Deadline:  time.Now().Add(-time.Hour),
			Tags:      []string{"tag1", "tag2"},
			Completed: false,
		}

		err := pastDeadlineTask.ValidateTask()
		expectedErr := "deadline cannot be in past date"
		if err == nil {
			t.Error("Expected an error, but got no error")
		} else if err.Error() != expectedErr {
			t.Errorf("Expected error message: %s, but got: %s", expectedErr, err.Error())
		}
	})
}
