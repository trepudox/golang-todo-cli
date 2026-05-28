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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Long: `Mark a task as completed in your todo list.

Usage examples:
  golang-todo-cli complete <task-id>   - Complete a specific task by ID`,
  	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("invalid usage: you must specify a task id to complete")
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

		_, err = repository.ChangeTaskStatusById(uint16(id), "completed")
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		listCmd.Run(listCmd, nil)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
