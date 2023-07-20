package main

import (
	"fmt"
	"github.com/MyOrg/ToDo/internal/task"
	"github.com/MyOrg/ToDo/internal/taskManager"
	"sort"
	"time"
)

func main() {
	taskManagerI := taskManager.NewTaskManager()

	// Sample task 1
	task1 := task.Task{
		Title:    "Task 1",
		Deadline: time.Now().Add(time.Hour * (1)),
		Content:  "Some random task 1",
		Tags:     []string{"tag1"},
	}

	// Add task in to-do
	taskID, err := taskManagerI.AddTask(task1)
	if err != nil {
		fmt.Println("Error:", err, "! failed to add task!")
		return
	}

	// Sample task 2
	task2 := task.Task{
		Title:    "Task 2",
		Deadline: time.Now().Add(time.Hour * (2)),
		Content:  "Some random task 2",
		Tags:     []string{"tag1"},
	}

	// Add task2 in to-do
	_, err = taskManagerI.AddTask(task2)
	if err != nil {
		fmt.Println("Error:", err, "! failed to add task!")
		return
	}

	// Sample task 3
	task3 := task.Task{
		Title:    "Task 3",
		Deadline: time.Now().Add(time.Hour * (3)),
		Content:  "Some random task 3",
		Tags:     []string{"tag1"},
	}

	// Add task3 in to-do
	_, err = taskManagerI.AddTask(task3)
	if err != nil {
		fmt.Println("Error:", err, "! failed to add task!")
		return
	}

	//Fetch task1 from to-do
	Task, err := taskManagerI.GetTask(taskID)
	if err != nil {
		fmt.Println("Error:", err, "! failed to get task!")
		return
	}
	fmt.Printf("Task: %s, Content: %s, Deadline: %s, Completed: %t", Task.Title, Task.Content, Task.Deadline.String(), Task.Completed)

	//Modify task 1
	Task.Title = "Modified Task 1"
	Task.Completed = true
	err = taskManagerI.ModifyTask(*Task)
	if err != nil {
		fmt.Println("Error:", err, "! failed to modify task!")
		return
	}
	fmt.Printf("Task: %s, Content: %s, Deadline: %s, Completed: %t\n", Task.Title, Task.Content, Task.Deadline.String(), Task.Completed)

	//Remove a task
	//err = taskManagerI.RemoveTask(taskID)
	//if err != nil {
	//	fmt.Println("Error:", err, "! failed to delete task!")
	//	return
	//}

	filterFunc := func(task task.Task) bool {
		return task.Deadline.After(time.Now())
	}

	sortFunc := func(tasks []task.Task) {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Deadline.Before(tasks[j].Deadline)
		})
	}

	fmt.Println()
	tasks := taskManagerI.ListTasks(filterFunc, sortFunc)
	for _, TaskObj := range tasks {
		fmt.Printf("Task: %s, Content: %s, Deadline: %s\n", TaskObj.Title, TaskObj.Content, TaskObj.Deadline.String())
	}

	fmt.Println()
	taskManagerI.GetStatistics(nil)

	//startTime := time.Now().Add(-time.Hour * 24)
	//taskManager.getStatistics(&startTime)

	fmt.Println()
	activityLog := taskManagerI.GetActivityLog(nil)
	for _, entry := range activityLog {
		fmt.Printf("Log_String : %s Log_Time : %s", entry.Log, entry.ActivityTS)
		fmt.Println()
	}
	return
}
