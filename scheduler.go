package main

import (
	"container/heap"
	"fmt"
)

type Scheduler struct {
	taskQueue PriorityQueue
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		taskQueue: make(PriorityQueue, 0),
	}
}

func (s *Scheduler) AddTask(task *Task) {
	heap.Push(&s.taskQueue, task)
}

func (s *Scheduler) Run() {
	for s.taskQueue.Len() > 0 {
		task := heap.Pop(&s.taskQueue).(*Task)
		fmt.Printf("Running task %d with priority %d and duration %d\n", task.ID, task.Priority, task.Duration)
	}
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
