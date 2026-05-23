/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"trepudox/golang-todo-cli/internal/task"
)

const jsonPath = "data.json"

func GetAllTasks() ([]task.Task, error) {
	// os.ReadFile, que le meu jsonfile em []byte
	jsonfile, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read jsonfile at '%s' location: %w", jsonPath, err)
	}

	// cria o slice e passa a referencia dele pro Unmarshal
	var tasks []task.Task
	if err := json.Unmarshal(jsonfile, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse the jsonfile: %w", err)
	}

	return tasks, nil
}

func AddTask(t task.Task) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %s", err.Error())
	}

	return append(tasks, t), nil
}

func RemoveTaskById(id int) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to remove task: %s", err.Error())
	}

	found := false
	for i, t := range tasks {
		if found {
			tasks[i-1] = tasks[i]
			continue
		}

		if t.Id == id {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("failed to remove task: no task with ID '%d'", id)
	}

	return tasks, nil
}

func ChangeTaskStatusById(id int, status string) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to change task status: %s", err.Error())
	}

	found := false
	for _, t := range tasks {
		if t.Id == id {
			t.Status = status
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("failed to chage task status: no task with ID '%d'", id)
	}

	return tasks, nil
}

func WriteData(tasks []task.Task) error {
	stringJson, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to parse json: %w", err)
	}

	if err := os.WriteFile(jsonPath, stringJson, 0o666); err != nil {
		return fmt.Errorf("failed to write json at path '%s'", jsonPath)
	}

	return nil
}
