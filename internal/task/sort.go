/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package task

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

var validArgs = []string{"id", "name", "priority", "due", "status"}

var sortFunctions = map[string]func([]Task, bool){
	"id":       func(tasks []Task, ascending bool) { sortById(tasks, ascending) },
	"name":     func(tasks []Task, ascending bool) { sortByName(tasks, ascending) },
	"priority": func(tasks []Task, ascending bool) { sortByPriority(tasks, ascending) },
	"due":      func(tasks []Task, ascending bool) { sortByDue(tasks, ascending) },
	"status":   func(tasks []Task, ascending bool) { sortByStatus(tasks, ascending) },
}

func SortTasks(tasks []Task, fieldName string, ascending bool) error {
	if fieldName == "" {
		return nil
	}

	if !slices.Contains(validArgs, fieldName) {
		return fmt.Errorf("warning: invalid argument for sorting: '%s'. valid arguments: %s",
			fieldName, strings.Join(validArgs, ", "))
	}

	callSortFunction := sortFunctions[strings.ToLower(fieldName)]
	if callSortFunction == nil {
		return fmt.Errorf("warning: no field found with name '%s'", fieldName)
	}

	callSortFunction(tasks, ascending)
	return nil
}

func sortById(tasks []Task, ascending bool) {
	slices.SortFunc(tasks, func(t1, t2 Task) int {
		if ascending {
			return int(t1.Id - t2.Id)
		}

		return int(t2.Id - t1.Id)
	})
}

func sortByName(tasks []Task, ascending bool) {
	slices.SortFunc(tasks, func(t1, t2 Task) int {
		if ascending {
			return strings.Compare(strings.ToLower(t1.Name), strings.ToLower(t2.Name))
		}

		return strings.Compare(strings.ToLower(t2.Name), strings.ToLower(t1.Name))
	})
}

func sortByPriority(tasks []Task, ascending bool) {
	slices.SortFunc(tasks, func(t1, t2 Task) int {
		if ascending {
			return comparePriority(t1.Priority, t2.Priority)
		}

		return comparePriority(t2.Priority, t1.Priority)
	})
}

func comparePriority(p1, p2 string) int {
	priorityWeight := map[string]int{
		"high":   3,
		"medium": 2,
		"low":    1,
	}

	return priorityWeight[p1] - priorityWeight[p2]
}

func sortByDue(tasks []Task, ascending bool) {
	slices.SortFunc(tasks, func(t1, t2 Task) int {
		t1Date, err := time.Parse("2006-01-02", t1.Due)
		if err != nil {
			fmt.Printf("failed to sort tasks by date. error: failed to parse due date of task ID '%d'", t1.Id)
			return 0
		}

		t2Date, err := time.Parse("2006-01-02", t2.Due)
		if err != nil {
			fmt.Printf("failed to sort tasks by date. error: failed to parse due date of task ID '%d'", t2.Id)
			return 0
		}

		if ascending {
			return t1Date.Compare(t2Date)
		}

		return t2Date.Compare(t1Date)
	})
}

func sortByStatus(tasks []Task, ascending bool) {
	slices.SortFunc(tasks, func(t1, t2 Task) int {
		if ascending {
			return compareStatus(t1.Status, t2.Status)
		}

		return compareStatus(t2.Status, t1.Status)
	})
}

func compareStatus(s1, s2 string) int {
	statusWeight := map[string]int{
		"completed": 2,
		"todo":      1,
	}

	return statusWeight[s1] - statusWeight[s2]
}
