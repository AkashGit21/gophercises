package cmd

import (
	"fmt"
	"os"

	"github.com/AkashGit21/gophercises/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "lists all of our incomplete tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := db.AllTasks()
			if err != nil {
				fmt.Println("Something went wrong:", err.Error())
				os.Exit(1)
			}
			if len(tasks) == 0 {
				fmt.Println("You have no task to work on. Why not take a break? ")
				return
			}

			fmt.Println("You have the following tasks: ")
			for ind, task := range tasks {
				fmt.Printf("%d. %s\n", ind+1, task.Value)
			}
		},
	}

	rootCmd.AddCommand(listCmd)
}
