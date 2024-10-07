package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	tm := NewTaskManager()

	for i := 0; i < 3; i++ {
		go tm.manageTasks()
	}

	for i := 0; i < 10; i++ {
		task := &Task{
			ID:       i,
			Priority: rand.Intn(5),
		}
		tm.AddTask(task)
		fmt.Printf("Added task %d with priority %d\n", task.ID, task.Priority)
		// time.Sleep(100 * time.Millisecond)
	}

	go func() {
		time.Sleep(time.Hour)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Program stopped.")
}
