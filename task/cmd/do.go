package cmd

import (
	"fmt"
	"os"
	"strconv"

	"example.com/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				color.New(color.FgRed).Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ReadAllTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		} else {
			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					color.New(color.FgRed).Println("Invalid task number:", id)
					continue
				}

				task := tasks[id-1]
				err := db.UpdateTaskCompleted(task.ID)

				if err != nil {
					fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err.Error())
				} else {
					color.New(color.FgGreen).Printf("Marked \"%d. %s\" as completed.\n", id, task.Description)
				}
			}
		}
	},
}
