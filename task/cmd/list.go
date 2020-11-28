package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"example.com/db"
)

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&taskType, "type", "t", "all", "Lists certain type of tasks.\n all: all tasks\n c: completed tasks\n nc: not completed tasks\n")
}

var taskType string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		taskType = strings.ToLower(taskType)
		switch taskType {
		case "all":
			printAllTasks()
		case "c":
			printSomeTasks(true)
		case "nc":
			printSomeTasks(false)
		default:
			fmt.Println("Unrecognized type. Please try again!")
		}
	},
}

func printAllTasks() {
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
}

func printSomeTasks(completed bool) {
	tasks, err := db.ReadSomeTasks(completed)

	if err != nil {
		fmt.Println("Something went wrong:", err.Error())
		os.Exit(1)
	}

	if completed {
		if len(tasks) == 0 {
			fmt.Println("You have not completed any tasks!")
		} else {
			fmt.Println("You have completed the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Description)
			}
		}
	} else {
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
		} else {
			fmt.Println("You have not completed the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Description)
			}
		}
	}
}
