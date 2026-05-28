/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/trepudox/golang-todo-cli/internal/repository"
)

// uncompleteCmd represents the uncomplete command
var uncompleteCmd = &cobra.Command{
	Use:   "uncomplete",
	Short: "Revert a completed task back to 'todo' status",
	Long: `Change the status of a completed task back to 'todo'.

This change the task status from 'completed' back to 'todo', allowing you to work on it again.

Usage examples:
  golang-todo-cli uncomplete <task-id>   - Revert a specific task status to 'todo'`,
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("invalid usage: you must specify a task id to uncomplete")
		}

		if _, err := strconv.ParseUint(args[0], 10, 16); err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		id, err := strconv.ParseUint(arg, 10, 16)
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = repository.ChangeTaskStatusById(uint16(id), "todo")
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		listCmd.Run(listCmd, nil)
	},
}

func init() {
	rootCmd.AddCommand(uncompleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uncompleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uncompleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
