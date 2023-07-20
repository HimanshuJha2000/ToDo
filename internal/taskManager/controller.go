package taskManager

import (
	"errors"
	"fmt"
	"github.com/MyOrg/ToDo/internal/task"
	"time"
)

// AddTask adds a new task to the TODO list
func (tm *TaskManager) AddTask(task task.Task) (int, error) {
	if err := task.ValidateTask(); err != nil {
		return 0, err
	}

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	task.ID = tm.NextTaskID
	tm.Tasks[task.ID] = task
	tm.NextTaskID++
	tm.ActivityLog = append(tm.ActivityLog, ActivityLogStruct{
		Log:        fmt.Sprintf("Added task: %s", task.Title),
		ActivityTS: time.Now(),
	})

	fmt.Println("Task added successfully!")
	fmt.Println("Task ID:", task.ID)
	return task.ID, nil
}

// GetTask retrieves a task by ID from the TODO list
func (tm *TaskManager) GetTask(taskID int) (*task.Task, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	Task, exists := tm.Tasks[taskID]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &Task, nil
}

// ModifyTask modifies an existing task in the TODO list
func (tm *TaskManager) ModifyTask(Task task.Task) error {
	if _, exists := tm.Tasks[Task.ID]; !exists {
		return errors.New(fmt.Sprintf("task with ID %d not found", Task.ID))
	}
	if err := Task.ValidateTask(); err != nil {
		return err
	}

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.Tasks[Task.ID] = Task
	tm.ActivityLog = append(tm.ActivityLog, ActivityLogStruct{
		Log:        fmt.Sprintf("Modified task: %s", Task.Title),
		ActivityTS: time.Now(),
	})
	fmt.Println("Task has been successfully modified!")
	return nil
}

// RemoveTask removes a task from the TODO list
func (tm *TaskManager) RemoveTask(taskID int) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	Task, exists := tm.Tasks[taskID]
	if !exists {
		return errors.New("task not found")
	}

	delete(tm.Tasks, taskID)
	tm.ActivityLog = append(tm.ActivityLog, ActivityLogStruct{
		Log:        fmt.Sprintf("Removed task: %s", Task.Title),
		ActivityTS: time.Now(),
	})

	fmt.Printf("Task with ID %d deleted successfully!", taskID)
	return nil
}

// ListTasks returns a list of tasks that match the given filter and sorted based on the defined sort criteria
func (tm *TaskManager) ListTasks(filterFunc func(task.Task) bool, sortFunc func([]task.Task)) []task.Task {
	filteredTasks := make([]task.Task, 0)
	for _, Task := range tm.Tasks {
		if filterFunc(Task) {
			filteredTasks = append(filteredTasks, Task)
		}
	}
	sortFunc(filteredTasks)
	return filteredTasks
}

// GetStatistics returns statistics for the given time period (optional)
func (tm *TaskManager) GetStatistics(timePeriod *time.Time) {
	var tasksAdded, tasksCompleted, tasksOverdue int

	for _, taskObj := range tm.Tasks {
		if timePeriod == nil || taskObj.Deadline.Before(*timePeriod) {
			tasksAdded++
			if taskObj.Completed {
				tasksCompleted++
			}
			if taskObj.Deadline.Before(time.Now()) && !taskObj.Completed {
				tasksOverdue++
			}
		}
	}

	fmt.Printf("Tasks added: %d\n", tasksAdded)
	fmt.Printf("Tasks completed: %d\n", tasksCompleted)
	fmt.Printf("Tasks overdue: %d\n", tasksOverdue)
}

// GetActivityLog returns the activity log for the given time period(optional)
func (tm *TaskManager) GetActivityLog(timePeriod *time.Time) []ActivityLogStruct {
	log := make([]ActivityLogStruct, 0)
	for _, entry := range tm.ActivityLog {
		if timePeriod == nil {
			log = append(log, entry)
		} else {
			if entry.ActivityTS.Before(*timePeriod) {
				log = append(log, entry)
			}
		}
	}
	return log
}
