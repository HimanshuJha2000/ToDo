package taskManager_test

import (
	"github.com/MyOrg/ToDo/internal/taskManager"
	"testing"
)

func TestNewTaskManager(t *testing.T) {
	t.Run("NewTaskManager non-nil", func(t *testing.T) {
		taskManagerI := taskManager.NewTaskManager()
		if taskManagerI == nil {
			t.Error("Expected a non-nil TaskManager, but got nil")
		}
	})

	t.Run("NewTaskManager Tasks initialized", func(t *testing.T) {
		taskManagerI := taskManager.NewTaskManager()

		// Check if the Tasks field is initialized as an empty map
		if taskManagerI.Tasks == nil {
			t.Error("Expected Tasks to be initialized as an empty map, but got nil")
		}
	})
}
