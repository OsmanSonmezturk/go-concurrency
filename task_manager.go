package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type TaskManager struct {
	tasks     chan *Task
	priorityQ PriorityQueue
	mu        sync.Mutex
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{
		tasks:     make(chan *Task, 100),
		priorityQ: make(PriorityQueue, 0),
	}
	heap.Init(&tm.priorityQ)
	return tm
}

func (tm *TaskManager) AddTask(task *Task) {
	tm.tasks <- task
}

func (tm *TaskManager) manageTasks() {
	for {
		select {
		case task := <-tm.tasks:
			tm.mu.Lock()
			heap.Push(&tm.priorityQ, task)
			tm.mu.Unlock()
		default:
			tm.mu.Lock()
			if tm.priorityQ.Len() > 0 {
				// tm.printPriorityQueue()
				task := heap.Pop(&tm.priorityQ).(*Task)
				tm.mu.Unlock()
				fmt.Printf("task started id: %d prio: %d \n", task.ID, task.Priority)

				time.Sleep(3000 * time.Millisecond)
			} else {
				tm.mu.Unlock()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (tm *TaskManager) printPriorityQueue() {
	if tm.priorityQ == nil {
		fmt.Println("  Queue is nil")
		return
	}
	for i, task := range tm.priorityQ {
		if task != nil {
			fmt.Printf("  %d: Task ID %d, Priority %d\n", i, task.ID, task.Priority)
		} else {
			fmt.Printf("  %d: nil task\n", i)
		}
	}
	fmt.Println()
}
