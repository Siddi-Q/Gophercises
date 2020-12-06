package cmd

import (
	"fmt"
	"os"
	"strings"

	"example.com/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		} else {
			color.New(color.FgGreen).Printf("Added \"%s\" to your task list.\n", task)
		}
	},
}
