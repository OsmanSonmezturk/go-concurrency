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
	wg        sync.WaitGroup
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
	tm.wg.Add(1)
	tm.tasks <- task
}

func (tm *TaskManager) manageTasks() {
	for {
		select {
		case task := <-tm.tasks:
			heap.Push(&tm.priorityQ, task)
		default:
			if tm.priorityQ.Len() > 0 {
				task := heap.Pop(&tm.priorityQ).(*Task)
				go tm.processTask(task)
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (tm *TaskManager) processTask(t *Task) {
	fmt.Printf("task started id: %d prio: %d \n", t.ID, t.Priority)
	time.Sleep(1000 * time.Millisecond) // task process
	tm.wg.Done()
}
