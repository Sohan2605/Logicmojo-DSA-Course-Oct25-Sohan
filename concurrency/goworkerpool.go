package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done() // marks this worker done when finished
    for job := range jobs {
        result := job * job
        fmt.Println("Worker", id, "processed job", job)
        results <- result
    }
    fmt.Println("Worker", id, "exiting")
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    var wg sync.WaitGroup
    numWorkers := 3

    // start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // no more jobs

    // wait for all workers to finish
    go func() {
        wg.Wait()
        close(results) // close results after all workers done
    }()

    // read results
    for res := range results {
        fmt.Println("Result:", res)
    }

    fmt.Println("All jobs processed")
}
