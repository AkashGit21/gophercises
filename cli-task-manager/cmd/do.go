package cmd

import (
	"fmt"
	"strconv"

	"github.com/AkashGit21/gophercises/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
	doCmd := &cobra.Command{
		Use:   "do",
		Short: "marks a task as complete",
		Run: func(cmd *cobra.Command, args []string) {
			var ids []int
			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("Failed to parse the argument:", arg)
					// return
				} else {
					ids = append(ids, id)
				}
			}

			tasks, _ := db.AllTasks()
			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					fmt.Println("invalid task number:", id)
					continue
				}
				task := tasks[id-1]
				err := db.DeleteTask(task.Key)
				if err!=nil {
					fmt.Printf("Failed to mark task \"%d\" as cmpleted. Error: %v\n", id,err)
				} else {
					fmt.Printf("Marked \"%d\" as completed.", id)
				}
			}

		},
	}

	rootCmd.AddCommand(doCmd)
}
