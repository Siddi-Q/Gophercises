package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"example.com/db"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadAllTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
		} else {
			fmt.Println("You have the following tasks:")
			for i, task := range tasks {
				var status string
				if task.Completed {
					status = "Completed"
				} else {
					status = "Not Completed"
				}

				fmt.Printf("%d. %s | %s\n", i+1, task.Description, status)
			}
		}
	},
}
