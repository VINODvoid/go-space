package main

import "fmt"

type Todo struct {
	ID    int
	Title string
	Done  bool
}
type TodoList []Todo

func (t *TodoList) Add(title string) {
	id := len(*t)
	todo := Todo{
		ID:    id + 1,
		Title: title,
		Done:  false,
	}
	*t = append(*t, todo)
}

func (t *TodoList) List() {
	for _, v := range *t {
		status := v.Done
		if status {
			fmt.Printf("[%v] %v (done) \n", v.ID, v.Title)
		} else {
			fmt.Printf("[%v] %v (pending) \n", v.ID, v.Title)
		}
	}
}

func (t *TodoList) Delete(id int) {
	for i, v := range *t {
		if v.ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			fmt.Println("Deleted the todo", v.ID)
			return
		}
	}
}

func (t *TodoList) Done(id int) {
	for i, v := range *t {
		if v.ID == id {
			(*t)[i].Done = true
			return
		}
	}
}
