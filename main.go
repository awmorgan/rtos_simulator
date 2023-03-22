package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("RTOS Simulator")

	// Command line flags
	addTask := flag.NewFlagSet("add", flag.ExitOnError)
	taskPriority := addTask.Int("p", 0, "Task priority")
	taskDuration := addTask.Int("d", 0, "Task duration")

	// Create a new scheduler
	scheduler := NewScheduler()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			fmt.Println("Invalid command. Use 'add' to add a task, 'run' to start the scheduler, or 'quit' to exit.")
			continue
		}

		switch args[0] {
		case "add":
			addTask.Parse(args[1:])
			if *taskPriority <= 0 || *taskDuration <= 0 {
				fmt.Println("Invalid task priority or duration. Must be greater than 0.")
				continue
			}
			taskID := scheduler.taskQueue.Len() + 1
			task := NewTask(taskID, *taskPriority, *taskDuration)
			scheduler.AddTask(task)
			fmt.Printf("Added task %d with priority %d and duration %d\n", task.ID, task.Priority, task.Duration)
		case "run":
			scheduler.Run()
		case "quit":
			fmt.Println("Exiting RTOS Simulator")
			return
		default:
			fmt.Println("Invalid command. Use 'add' to add a task, 'run' to start the scheduler, or 'quit' to exit.")
		}
	}
}
