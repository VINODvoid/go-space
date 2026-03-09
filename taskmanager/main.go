package main

import (
	"flag"
	"fmt"
)

func main() {
	tasks := Store{File: "tasks.json"}
	tasks.Load()
	add := flag.String("add", "", "add your task")
	list := flag.Bool("list", false, "list your tasks")
	done := flag.Int("done", 0, "mark your task done")
	flag.Parse()
	if *add != "" {
		tasks.Add(*add)
		tasks.Save()
	}
	if *list {
		collection := tasks.List()
		for _, value := range collection {
			fmt.Printf("%s ( %v )\n", value.Title, value.Done)
		}
	}
	if *done != 0 {
		tasks.Done(*done)
		tasks.Save()
	}
}
