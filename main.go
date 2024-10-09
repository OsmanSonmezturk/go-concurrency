package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tm := NewTaskManager()
	go tm.manageTasks()

	for i := 0; i < 10; i++ {
		task := &Task{
			ID:       i,
			Priority: rand.Intn(5),
		}
		tm.AddTask(task)
		fmt.Printf("Added task %d with priority %d\n", task.ID, task.Priority)

	}

	tm.wg.Wait()
	close(tm.tasks)
	fmt.Println("Program stopped.")
}
