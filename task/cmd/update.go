package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task's description",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			color.New(color.FgRed).Println("Failed to parse the argument:", args[0])
			os.Exit(1)
		}

		description := strings.Join(args[1:], " ")

		tasks, err := db.ReadAllTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		} else {
			if id <= 0 || id > len(tasks) {
				color.New(color.FgRed).Println("Invalid task number:", id)
			} else {
				task := tasks[id-1]
				err := db.UpdateTaskDescription(task.ID, description)

				if err != nil {
					fmt.Printf("Failed to update \"%d. %s\". Error: %s\n", id, task.Description, err.Error())
				} else {
					color.New(color.FgGreen).Printf("Updated task's description from \"%d. %s\" to \"%d. %s\".\n", id, task.Description, id, description)
				}
			}
		}
	},
}
