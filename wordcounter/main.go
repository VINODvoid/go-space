package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Add filename in this format ==> go run main.go filename")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counts := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			counts[word]++
		}
	}
	sortedWords := []string{}
	for key := range counts {
		sortedWords = append(sortedWords, key)
	}
	sort.Strings(sortedWords)
	for _, word := range sortedWords {
		fmt.Println(word, counts[word])
	}
}
