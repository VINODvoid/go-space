package main

import (
	"fmt"
	"time"
)

func rateLimiter(rate time.Duration, jobs []string) {
	ticker := time.NewTicker(rate * time.Second)

	defer ticker.Stop()
	for _, job := range jobs {
		<-ticker.C

		fmt.Printf("%s -- %v\n", job, time.Now().Format("15:04:05"))
	}
}

func main() {
	jobs := []string{"job1", "job2", "job3", "job4", "job5"}
	rateLimiter(3, jobs)
}
