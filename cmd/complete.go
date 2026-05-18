/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Long: `Mark one or all tasks as completed in your todo list.

Usage examples:
  golang-todo-cli complete <task-id>       - Complete a specific task by ID
  golang-todo-cli complete --all           - Mark all pending tasks as done

Flags:
  --all Complete all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
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
