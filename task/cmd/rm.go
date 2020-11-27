package cmd

import (
	"fmt"
	"os"
	"strconv"

	"example.com/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
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
					fmt.Println("Invalid task number:", id)
					continue
				}

				task := tasks[id-1]
				err := db.DeleteTask(task.ID)

				if err != nil {
					fmt.Printf("Failed to delete \"%d\". Error: %s\n", id, err.Error())
				} else {
					fmt.Printf("Deleted \"%d. %s\".\n", id, task.Description)
				}
			}
		}
	},
}
