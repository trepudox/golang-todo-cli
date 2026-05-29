/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/trepudox/golang-todo-cli/internal/repository"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Update details of an existing task",
	Long: `Modify specific fields of a task identified by its ID.

You can update the task name, priority level, or due date independently 
or all at once. Only the provided flags will be updated.

Usage examples:
  todo edit <id> --name "New Task Name"
  todo edit <id> --priority high
  todo edit <id> --due-date 2026-12-31
  todo edit <id> --name "Updated" --priority low

Flags:
  --name string       Update the task title/description
  --priority string   Set priority (low, medium, high)
  --due-date string   Set due date in YYYY-MM-DD format`,
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("invalid usage: you must specify a task id to uncomplete")
		}

		if _, err := strconv.ParseUint(args[0], 10, 16); err != nil {
			return err
		}

		if newDueDate != "" {
			_, err := time.Parse("2006-01-02", newDueDate)
			if err != nil {
				return fmt.Errorf("invalid due date: value must be in the following format: 'YYYY-MM-DD'. error: %s\n", err.Error())
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		id, err := strconv.ParseUint(arg, 10, 16)
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = repository.EditTaskById(uint16(id), newName, newPriority, newDueDate)
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		listCmd.Run(listCmd, nil)
	},
}

var newName string
var newPriority string
var newDueDate string

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&newName, "name", "n", "", "The new name for the task")
	editCmd.Flags().StringVarP(&newPriority, "priority", "p", "", "The new priority for the task")
	editCmd.Flags().StringVarP(&newDueDate, "due-date", "d", "", "The new due date for the task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
