/*
Copyright © 2026 trepudox <teoaa2@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
