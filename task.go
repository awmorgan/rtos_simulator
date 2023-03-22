package main

type Task struct {
	ID       int
	Priority int
	Duration int
}

func NewTask(id, priority, duration int) *Task {
	return &Task{
		ID:       id,
		Priority: priority,
		Duration: duration,
	}
}
