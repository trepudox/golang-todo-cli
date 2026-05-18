/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
  --priority string   Set task priority (low, medium, high)
  --due      string   Set a due date for the task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
