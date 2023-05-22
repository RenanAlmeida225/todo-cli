package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Todo

func (t *Todos) loadFile() {
	file, err := os.ReadFile("./todos.json")
	if err != nil {
		fmt.Println("error file", err)
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		fmt.Println("error json", err)
	}
}

func (t *Todos) saveInFile() {
	data, err := json.Marshal(t)
	if err != nil {
		return
	}
	os.WriteFile("./todos.json", data, 0644)
}

func (t *Todos) Save(task string) {
	t.loadFile()
	todo := Todo{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
	t.saveInFile()
}

func (t *Todos) Complete(id int) {
	t.loadFile()
	todos := *t
	if id <= 0 || id > len(todos) {
		return
	}
	todos[id-1].Done = true
	todos[id-1].CompletedAt = time.Now()
	t.saveInFile()
}

func (t *Todos) Delete(id int) {
	t.loadFile()
	todos := *t
	if id <= 0 || id > len(todos) {
		return
	}
	*t = append(todos[:id-1], todos[id:]...)
	t.saveInFile()
}

func (t *Todos) List() {
	t.loadFile()
	todos := *t
	for i, todo := range todos {
		fmt.Printf("id %d | task %s | done %t |createAt %s | completeAt %s\n", (i + 1), todo.Task, todo.Done, todo.CreatedAt, todo.CompletedAt)
	}
}
