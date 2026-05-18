/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"
	"trepudox/golang-todo-cli/internal/repository"
	"trepudox/golang-todo-cli/internal/task"
	"trepudox/golang-todo-cli/internal/ui"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display tasks in your todo list",
	Long: `List all tasks from your todo list, with optional filtering and/or sorting.

Usage examples:
  golang-todo-cli list                      - Show all tasks
  golang-todo-cli list --status todo        - Show only pending tasks
  golang-todo-cli list --priority high      - Filter by priority level
  golang-todo-cli list --sort-asc name      - Sort ascendingly by name (from Z to A)
  golang-todo-cli list --sort-desc priority - Sort descendingly by priority level (from high to low)

Flags:
  --status    string    Filter by status (todo, completed)
  --priority  string    Filter by priority level (low, medium, high)
  --sort-asc  string    Sort tasks in ascending order by the specified field
  --sort-desc string    Sort tasks in descending order by the specified field`,
	Args: cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("sort-desc") {
			ascendingSort = false
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := repository.ReadData()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// filter pela task
		tasks, err = task.FilterTasksByStatus(tasks, statusToFilter)
		if err != nil {
			fmt.Printf("%s\nproceeding without filtering\n\n", err.Error())
		}

		tasks, err = task.FilterTasksByPriority(tasks, priorityToFilter)
		if err != nil {
			fmt.Printf("%s\nproceeding without filtering\n\n", err.Error())
		}

		// sort pela task
		if err := task.SortTasks(tasks, fieldNameToSort, ascendingSort); err != nil {
			fmt.Printf("%s\nproceeding without sorting\n\n", err.Error())
		}

		if err := ui.ListTasks(tasks); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

var statusToFilter string
var priorityToFilter string
var fieldNameToSort string
var ascendingSort = true

func init() {
	rootCmd.AddCommand(listCmd)

	// filter
	listCmd.Flags().StringVarP(&statusToFilter, "status", "s", "", "Filter by status (todo, completed)")
	listCmd.Flags().StringVarP(&priorityToFilter, "priority", "p", "", "Filter by priority level (low, medium, high)")

	// sort
	listCmd.Flags().StringVar(&fieldNameToSort, "sort-asc", "", "Sort tasks in ascending order by the specified field")
	listCmd.Flags().StringVar(&fieldNameToSort, "sort-desc", "", "Sort tasks in descending order by the specified field")

	listCmd.MarkFlagsMutuallyExclusive("sort-asc", "sort-desc")
}
