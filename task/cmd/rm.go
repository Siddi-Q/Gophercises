package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rm")
	},
}
