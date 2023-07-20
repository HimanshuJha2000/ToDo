package taskManager_test

import (
	"github.com/MyOrg/ToDo/internal/task"
	"github.com/MyOrg/ToDo/internal/taskManager"
	"testing"
	"time"
)

func mockTaskManager() *taskManager.TaskManager {
	return &taskManager.TaskManager{
		Tasks:       make(map[int]task.Task),
		NextTaskID:  1,
		ActivityLog: make([]taskManager.ActivityLogStruct, 0),
	}
}

func TestAddTask(t *testing.T) {
	t.Run("ValidTask", func(t *testing.T) {
		sampleTask := task.Task{
			ID:       1,
			Title:    "Test Task",
			Deadline: time.Now().Add(time.Hour),
		}

		mockTaskManagerObj := mockTaskManager()
		_, err := mockTaskManagerObj.AddTask(sampleTask)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("InvalidTask", func(t *testing.T) {
		mockInvalidTask := task.Task{
			ID:    0,
			Title: "",
		}
		mockTaskManagerObj := mockTaskManager()

		_, err := mockTaskManagerObj.AddTask(mockInvalidTask)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestGetTask(t *testing.T) {
	t.Run("ExistingTask", func(t *testing.T) {
		mockTaskManagerObj := mockTaskManager()
		sampleTask := task.Task{
			ID:       1,
			Title:    "Test Task",
			Deadline: time.Now().Add(time.Hour),
		}
		mockTaskManagerObj.Tasks[sampleTask.ID] = sampleTask

		Task, _ := mockTaskManagerObj.GetTask(sampleTask.ID)
		if Task.ID != sampleTask.ID {
			t.Errorf("Expected task %v, got %v", sampleTask, Task.ID)
		}
	})

	t.Run("NonExistingTask", func(t *testing.T) {
		mockTaskManagerObj := mockTaskManager()

		_, err := mockTaskManagerObj.GetTask(1)
		expectedError := "task not found"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
		}
	})
}

func TestModifyTask(t *testing.T) {
	t.Run("ExistingTask_ValidModification", func(t *testing.T) {
		mockTaskManagerObj := mockTaskManager()
		sampleTask := task.Task{
			ID:       1,
			Title:    "Test Task",
			Deadline: time.Now().Add(time.Hour),
		}
		mockTaskManagerObj.Tasks[sampleTask.ID] = sampleTask
		modifiedTask := task.Task{
			ID:       sampleTask.ID,
			Title:    "Modified Test Task",
			Deadline: time.Now().Add(time.Hour),
		}
		err := mockTaskManagerObj.ModifyTask(modifiedTask)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("ExistingTask_InvalidModification", func(t *testing.T) {
		mockTaskManagerObj := mockTaskManager()
		sampleTask := task.Task{
			ID:       1,
			Title:    "Test Task",
			Deadline: time.Now().Add(time.Hour),
		}
		mockTaskManagerObj.Tasks[sampleTask.ID] = sampleTask
		invalidTask := task.Task{
			ID:    sampleTask.ID,
			Title: "",
		}
		err := mockTaskManagerObj.ModifyTask(invalidTask)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
