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

func ReadData() ([]task.Task, error) {
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
