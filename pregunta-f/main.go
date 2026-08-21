package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID int
}

type Runner struct{}

func (Runner) Execute(t Task, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()
	out <- t.ID * 10
}

func main() {
	out := make(chan int, 3)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		Runner{}.Execute(Task{ID: i}, &wg, out)
	}
	wg.Wait()
	close(out)
	for v := range out {
		fmt.Println(v)
	}
}
