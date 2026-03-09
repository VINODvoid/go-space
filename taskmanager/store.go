package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Store struct {
	Tasks []Task `json:"tasks"`
	File  string `json:"file"`
}

func (store *Store) Load() error {
	fileOpen, err := os.Open(store.File)
	if err != nil {
		return nil
	}
	defer fileOpen.Close()
	errr := json.NewDecoder(fileOpen).Decode(&store.Tasks)
	if errr != nil {
		return fmt.Errorf("Unable to load tasks")
	}

	return nil
}

func (store *Store) Save() error {
	fileOpen, err := os.Create(store.File)
	if err != nil {
		return fmt.Errorf("%s is not present in the directory", store.File)
	}
	defer fileOpen.Close()
	errr := json.NewEncoder(fileOpen).Encode(store.Tasks)
	if errr != nil {
		return fmt.Errorf("Unable to load tasks")
	}
	return nil
}

func (store *Store) Add(title string) error {
	id := len(store.Tasks)

	task := Task{
		ID:    id + 1,
		Title: title,
		Done:  false,
	}
	store.Tasks = append(store.Tasks, task)
	return nil
}

func (store *Store) Delete(id int) error {
	for i, v := range store.Tasks {
		if v.ID == id {
			store.Tasks = append((store.Tasks)[:i], (store.Tasks)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (store *Store) Done(id int) error {
	for i, v := range store.Tasks {
		if v.ID == id {
			store.Tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (store *Store) List() []Task {
	return store.Tasks
}
