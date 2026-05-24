/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package task

import (
	"fmt"
	"slices"
	"strings"
)

var priorityValidArgs = []string{"low", "medium", "high"}
var statusValidArgs = []string{"todo", "completed"}

// sim, o codigo está duplicado!! melhor do que abstrair pra 2 funcoes nesse momento

func FilterTasksByPriority(tasks []Task, priority string) ([]Task, error) {
	if priority == "" {
		return tasks, nil
	}

	if !slices.Contains(priorityValidArgs, priority) {
		return tasks, fmt.Errorf("warning: invalid argument for filtering: '%s'. valid arguments: %s",
			priority, strings.Join(priorityValidArgs, ", "))
	}

	filteredSlice := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if priority == t.Priority {
			filteredSlice = append(filteredSlice, t)
		}
	}

	return filteredSlice, nil
}

func FilterTasksByStatus(tasks []Task, status string) ([]Task, error) {
	if status == "" {
		return tasks, nil
	}

	if !slices.Contains(statusValidArgs, status) {
		return tasks, fmt.Errorf("warning: invalid argument for filtering: '%s'. valid arguments: %s",
			status, strings.Join(statusValidArgs, ", "))
	}

	filteredSlice := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if status == t.Status {
			filteredSlice = append(filteredSlice, t)
		}
	}

	return filteredSlice, nil
}
