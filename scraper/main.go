package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Result struct {
	URL       string
	Status    int
	WordCount int
	Err       error
}

func fetch(url string) Result {
	resp, err := http.Get(url)
	if err != nil {
		return Result{Err: fmt.Errorf("URL is invalid"), URL: url}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Err: fmt.Errorf("Body cannot be parsed")}
	}
	words := strings.Fields(string(body))
	count := len(words)
	return Result{URL: url, Status: resp.StatusCode, WordCount: count}
}

func main() {
	url := os.Args[1:]
	chn := make(chan Result)
	var wg sync.WaitGroup
	for _, u := range url {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			chn <- fetch(u)
		}(u)
	}
	go func() {
		wg.Wait()
		close(chn)
	}()
	for result := range chn {
		if result.Err != nil {
			fmt.Printf("ERROR %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("%d %s — %d words\n", result.Status, result.URL, result.WordCount)
		}
	}
}
