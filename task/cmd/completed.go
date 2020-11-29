package cmd

import (
	"fmt"
	"os"

	"example.com/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists all of your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadSomeTasks(true)

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have not completed any tasks!")
		} else {
			fmt.Println("You have completed the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s | Completed on %v\n", i+1, task.Description, task.CompletedDate.Format("Monday, January 2, 2006 03:04:05 PM"))
			}
		}
	},
}
