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
	RootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&rmAll, "all", "a", false, "Removes all tasks\n")
}

var rmAll bool

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		if rmAll {
			db.DeleteBucket()
			color.New(color.FgYellow).Println("All tasks were deleted.")
		} else {
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
					err := db.DeleteTask(task.ID)

					if err != nil {
						fmt.Printf("Failed to delete \"%d\". Error: %s\n", id, err.Error())
					} else {
						color.New(color.FgYellow).Printf("Deleted \"%d. %s\".\n", id, task.Description)
					}
				}
			}
		}
	},
}
