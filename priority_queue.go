package main

type Task struct {
	ID       int
	Priority int
	Data     interface{}
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	result := pq[i].Priority < pq[j].Priority
	//fmt.Printf("in less function")
	return result
}

func (pq PriorityQueue) Swap(i, j int) {
	// fmt.Printf("in swap function")
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
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
