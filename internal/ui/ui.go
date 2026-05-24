/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package ui

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/trepudox/golang-todo-cli/internal/task"
)

const minWidth int = 0
const tabWidth int = 8
const padding int = 4
const padChar byte = ' '

var w = tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padChar, tabwriter.StripEscape)

func printHeader(w *tabwriter.Writer) {
	fmt.Fprintln(w, "ID\tNAME\tPRIORITY\tDUE\tSTATUS")
}

func printTasks(w *tabwriter.Writer, data []task.Task) error {
	for _, t := range data {
		formattedDue, err := task.FormatDueDate(t.Due)
		if err != nil {
			return fmt.Errorf("failed to format date: %w", err)
		}

		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n",
			t.Id, t.Name, t.Priority, formattedDue, t.Status)
	}

	return nil
}

func ListTasks(data []task.Task) error {
	printHeader(w)

	if err := printTasks(w, data); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("failed to flush tabwriter: %w", err)
	}

	return nil
}
