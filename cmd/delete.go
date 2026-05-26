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

// TODO: criar a logica do delete all com uma confirmacao
// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a task from your todo list",
	Long: `Remove a specific task from your todo list.

Warning: This action cannot be undone. The task will be deleted completely.

Usage examples:
  golang-todo-cli delete <task-id>   - Delete a specific task by ID
  golang-todo-cli delete --all       - Remove all tasks (use with caution)

Flags:
  --all Delete all tasks`,
	Args: cobra.MaximumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if deleteAll && len(args) > 0 {
			return fmt.Errorf("invalid usage: you can only specify the task id or the '--all' flag")
		}

		if !deleteAll && len(args) == 0 {
			return fmt.Errorf("invalid usage: you must specify a task id to delete")
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

		_, err = repository.RemoveTaskById(uint16(id))
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		listCmd.Run(listCmd, nil)
	},
}

var deleteAll bool

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVarP(&deleteAll, "all", "a", false, "Remove all tasks (use with caution)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
