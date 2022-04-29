package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AkashGit21/gophercises/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the list of tasks",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			_, err := db.CreateTask(task)
			if err != nil {
				fmt.Println("Something went wrong:", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Added \"%s\" to your task list!\n", task)
		},
	}

	rootCmd.AddCommand(addCmd)
}
