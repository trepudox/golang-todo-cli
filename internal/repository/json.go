/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/trepudox/golang-todo-cli/internal/task"
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
	if err = json.Unmarshal(jsonfile, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse the jsonfile: %w", err)
	}

	return tasks, nil
}

func AddTask(t task.Task) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %s", err.Error())
	}

	if err = assignNewTaskId(&t); err != nil {
		return nil, fmt.Errorf("failed to assign an id to new task: %s", err.Error())
	}

	tasks = append(tasks, t)
	if err = writeData(tasks); err != nil {
		return nil, fmt.Errorf("failed to save data: %s", err.Error())
	}

	return tasks, nil
}

func RemoveTaskById(id uint16) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to remove task: %s", err.Error())
	}

	tasksOriginalLen := len(tasks)
	tasks = slices.DeleteFunc(tasks, func(t task.Task) bool {
		return t.Id == id
	})

	if tasksOriginalLen == len(tasks) {
		return nil, fmt.Errorf("failed to remove task: no task with ID '%d'", id)
	}

	if err = writeData(tasks); err != nil {
		return nil, fmt.Errorf("failed to save data: %s", err.Error())
	}

	return tasks, nil
}

func ChangeTaskStatusById(id uint16, status string) ([]task.Task, error) {
	tasks, err := GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to change task status: %s", err.Error())
	}

	found := false
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = status
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("failed to change task status: no task with ID '%d'", id)
	}

	if err = writeData(tasks); err != nil {
		return nil, fmt.Errorf("failed to save data: %s", err.Error())
	}

	return tasks, nil
}

func assignNewTaskId(t *task.Task) error {
	tasks, err := GetAllTasks()
	if err != nil {
		return err
	}

	var biggestId uint16 = 0
	for _, tt := range tasks {
		if tt.Id > biggestId {
			biggestId = tt.Id
		}
	}

	t.Id = biggestId + 1
	return nil
}

func writeData(tasks []task.Task) error {
	stringJson, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to parse json: %w", err)
	}

	if err = os.WriteFile(jsonPath, stringJson, 0o666); err != nil {
		return fmt.Errorf("failed to write json at path '%s'", jsonPath)
	}

	return nil
}
