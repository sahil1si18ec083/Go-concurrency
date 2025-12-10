package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobsList := []string{
		"image_1.png",
		"image_2.png",
		"image_3.png",
		"image_4.png",
		"image_5.png",
	}

	numWorkers := 3
	jobs := make(chan string, len(jobsList))
	results := make(chan Result, len(jobsList))

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	start := time.Now()

	for i := 1; i <= numWorkers; i++ {
		go Worker(i, jobs, results, &wg)
	}

	for _, job := range jobsList {
		jobs <- job
	}
	close(jobs)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println("Processed:", res.Value)
	}

	fmt.Println("Time taken:", time.Since(start))
}
