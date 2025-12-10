package main

import (
	"sync"
	"time"
)

type Result struct {
	Value string
	Err   error
}

func Worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Millisecond * 200)
		results <- Result{Value: job, Err: nil}
	}
}
