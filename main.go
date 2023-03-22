package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("RTOS Simulator")

	// Command line flags
	addTask := flag.NewFlagSet("add", flag.ExitOnError)
	taskPriority := addTask.Int("p", 0, "Task priority")
	taskDuration := addTask.Int("d", 0, "Task duration")

	// Create a new scheduler
	scheduler := NewScheduler()

	if len(os.Args) < 2 {
		fmt.Println("Invalid command. Use 'add' to add a task or 'run' to start the scheduler.")
		return
	}

	switch os.Args[1] {
	case "add":
		addTask.Parse(os.Args[2:])
		if *taskPriority <= 0 || *taskDuration <= 0 {
			fmt.Println("Invalid task priority or duration. Must be greater than 0.")
			return
		}
		taskID := scheduler.taskQueue.Len() + 1
		task := NewTask(taskID, *taskPriority, *taskDuration)
		scheduler.AddTask(task)
		fmt.Printf("Added task %d with priority %d and duration %d\n", task.ID, task.Priority, task.Duration)
	case "run":
		scheduler.Run()
	default:
		fmt.Println("Invalid command. Use 'add' to add a task or 'run' to start the scheduler.")
	}
}
