package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}
