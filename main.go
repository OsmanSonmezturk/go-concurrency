package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	tm := NewTaskManager()
	go tm.manageTasks()

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal, closing task channel")
		tm.FinishAddingTasks()
	}()

	for i := 0; i < 10; i++ {
		task := &Task{
			ID:       i,
			Priority: rand.Intn(5),
		}
		tm.AddTask(task)
		fmt.Printf("Added task %d with priority %d\n", task.ID, task.Priority)
	}

	tm.FinishAddingTasks()
	fmt.Println("Program stopped.")
}
