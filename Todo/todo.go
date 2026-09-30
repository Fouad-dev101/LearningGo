package main

import (
	"errors"
	"fmt"
	"time"
)

// Task represents a single to-do item.
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

// TodoList holds tasks in memory.
type TodoList struct {
	tasks  []Task
	nextID int
}

// NewTodoList creates an empty list.
func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  []Task{},
		nextID: 1,
	}
}

// Add inserts a new task and returns it.
func (t *TodoList) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	task := Task{
		ID:        t.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	t.nextID++
	t.tasks = append(t.tasks, task)
	return task, nil
}

// List returns all tasks.
func (t *TodoList) List() []Task {
	return t.tasks
}

// Complete marks a task as done by ID.
func (t *TodoList) Complete(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Delete removes a task by ID.
func (t *TodoList) Delete(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}