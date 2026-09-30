package main

import "testing"

func TestAdd(t *testing.T) {
	list := NewTodoList()

	task, err := list.Add("Buy milk")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "Buy milk" {
		t.Errorf("expected title 'Buy milk', got %q", task.Title)
	}
	if task.Done {
		t.Error("new task should not be done")
	}

	if _, err := list.Add(""); err == nil {
		t.Error("expected error for empty title, got nil")
	}
}

func TestComplete(t *testing.T) {
	list := NewTodoList()
	task, _ := list.Add("Walk dog")

	if err := list.Complete(task.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !list.List()[0].Done {
		t.Error("task should be marked done")
	}

	if err := list.Complete(999); err == nil {
		t.Error("expected error for missing ID, got nil")
	}
}

func TestDelete(t *testing.T) {
	list := NewTodoList()
	t1, _ := list.Add("Task A")
	_, _ = list.Add("Task B")

	if err := list.Delete(t1.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list.List()) != 1 {
		t.Errorf("expected 1 task remaining, got %d", len(list.List()))
	}
	if list.List()[0].Title != "Task B" {
		t.Errorf("wrong task remaining: %s", list.List()[0].Title)
	}

	if err := list.Delete(999); err == nil {
		t.Error("expected error for missing ID, got nil")
	}
}