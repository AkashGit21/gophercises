package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "lists all of our incomplete tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running list command!")
		},
	}

	rootCmd.AddCommand(listCmd)
}
