/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/trepudox/golang-todo-cli/internal/repository"
	"github.com/trepudox/golang-todo-cli/internal/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your todo list",
	Long: `Add a new task to your todo list with optional details.

Usage examples:
  golang-todo-cli add "Buy groceries"                    - Add a simple task
  golang-todo-cli add "Meeting at 3pm" --priority high   - Add with priority
  golang-todo-cli add "Finish report" --due 2026-12-31   - Add with due date

Flags:
  --priority string   Set task priority (low, medium, high) (default "medium")
  --due      string   Set a due date for the task (default "today")`,
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		switch priority {
		case "low", "medium", "high":
			break
		default:
			return fmt.Errorf("invalid priority: value must be 'low', 'medium' or 'high'\n")
		}

		_, err := time.Parse("2006-01-02", due)
		if err != nil {
			return fmt.Errorf("invalid due date: value must be in the following format: 'YYYY-MM-DD'. error: %s\n", err.Error())
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		newTask := task.NewTask(taskName, priority, due, "todo")

		_, err := repository.AddTask(newTask)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		fmt.Printf("task '%s' added sucessfully\n\n", taskName)

		listCmd.Run(listCmd, nil)
	},
}

var priority string
var due string

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Add with priority (low, medium or high)")
	addCmd.Flags().StringVarP(&due, "due", "d", time.Now().Format("2006-01-02"), "Add with due date (example: 2026-05-24)")
}
