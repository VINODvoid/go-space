package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	todos := TodoList{}
	todos.Add("buy milk")
	todos.Add("write code")
	todos.Add("go for a walk")
	todos.Done(1)
	method := os.Args[1]
	switch method {
	case "add":
		title := os.Args[2]
		todos.Add(title)
		fmt.Println("Todo is Added")
	case "delete":
		id := os.Args[2]
		todoId, _ := strconv.Atoi(id)

		todos.Delete(todoId)
	case "list":
		todos.List()
	case "done":
		id := os.Args[2]
		todoId, _ := strconv.Atoi(id)
		todos.Done(todoId)
	default:
		fmt.Println("Unable to perform any operation , check the arguments")
	}
}
