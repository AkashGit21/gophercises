package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the list of tasks",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			fmt.Printf("Added \"%s\" to your task list!\n", task)
		},
	}

	rootCmd.AddCommand(addCmd)
}
