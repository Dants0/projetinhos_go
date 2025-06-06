package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Tasks struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const tasksFile = "tasks.json"

var tasks []Tasks

func loadTasks() {
	data, err := ioutil.ReadFile(tasksFile)
	if err != nil {
		tasks = []Tasks{}
		return
	}

	_ = json.Unmarshal(data, &tasks)
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", " ")
	_ = ioutil.WriteFile(tasksFile, data, 0644)
}

func AddTask(titleTask string) {
	loadTasks()
	id := len(tasks) + 1
	newTask := Tasks{
		ID:        id,
		Title:     titleTask,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	saveTasks()
	fmt.Printf("Tarefa adicionada: [%d] %s\n", newTask.ID, newTask.Title)
}

func ListTasks() {
	loadTasks()
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa adicionada!")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("[%d] [%s] %s\n", task.ID, status, task.Title)
	}
}

func MarkDone(idStr string) {
	loadTasks()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if found {
		saveTasks()
		fmt.Printf("Tarefa %d concluída!", id)
	} else {
		fmt.Printf("Tarefa %d não encontrada!", id)
	}
}

func MarkUndone(idStr string) {
	loadTasks()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = false
			found = true
			break
		}
	}

	if found {
		saveTasks()
		fmt.Printf("Tarefa %d desconcluída!", id)
	} else {
		fmt.Printf("Tarefa %d não encontrada!", id)
	}
}

func DeleteTask(idStr string) {
	loadTasks()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	idx := -1
	for i, task := range tasks {
		if task.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Printf("Tarefa %d não encontrada. \n", id)
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)
	saveTasks()
	fmt.Printf("Tarefa %d excluída.\n", id)
}
