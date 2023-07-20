package taskManager

import (
	"github.com/MyOrg/ToDo/internal/task"
	"sync"
	"time"
)

// TaskManager represents the TODO task manager
type TaskManager struct {
	Tasks       map[int]task.Task
	NextTaskID  int
	ActivityLog []ActivityLogStruct
	mutex       sync.RWMutex
}

type ActivityLogStruct struct {
	Log        string
	ActivityTS time.Time
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:       make(map[int]task.Task),
		NextTaskID:  1,
		ActivityLog: make([]ActivityLogStruct, 0),
	}
}
