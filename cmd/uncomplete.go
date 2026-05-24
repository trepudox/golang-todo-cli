/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// uncompleteCmd represents the uncomplete command
var uncompleteCmd = &cobra.Command{
	Use:   "uncomplete",
	Short: "Revert a completed task back to 'todo' status",
	Long: `Change the status of a completed task back to 'todo'.

This change the task status from 'completed' back to 'todo', allowing you to work on it again.

Usage examples:
  golang-todo-cli uncomplete <task-id>   - Revert a specific task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uncomplete called")
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
