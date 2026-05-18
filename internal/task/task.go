/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package task

import (
	"fmt"
	"time"
)

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Priority string `json:"priority"`
	Due      string `json:"due"`
	Status   string `json:"status"`
}

func FormatDueDate(taskDue string) (string, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return "", fmt.Errorf("failed to load timezone: %w", err)
	}

	taskDueDate, err := time.Parse("2006-01-02", taskDue)
	if err != nil {
		return "", fmt.Errorf("failed to parse dueDate: %w", err)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	targetDate := time.Date(taskDueDate.Year(), taskDueDate.Month(), taskDueDate.Day(), 0, 0, 0, 0, loc)

	hoursDiff := targetDate.Sub(today).Hours()
	daysDiff := int(hoursDiff / 24)

	switch {
	case daysDiff == 0:
		return "today", nil
	case daysDiff == 1:
		return "tomorrow", nil
	case daysDiff == 7:
		return "a week from now", nil
	case daysDiff == 30:
		return "a month from now", nil
	case daysDiff == -1:
		return "yesterday", nil
	case daysDiff == -7:
		return "one week late", nil
	case daysDiff == -30:
		return "one month late", nil
	default:
		if daysDiff > 0 {
			return fmt.Sprintf("%d days from now", daysDiff), nil
		}

		return fmt.Sprintf("%d days late", -daysDiff), nil
	}

}
