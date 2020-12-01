package cmd

import (
	"fmt"
	"os"
	"time"

	"example.com/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(completedCmd)
	completedCmd.Flags().StringVarP(&duration, "duration", "d", "forever", "the time during which the tasks were completed\n forever: all tasks regardless of completed date\n today: all tasks completed today\n 24h: all tasks completed in the last 24 hours\n 12h: all tasks completed in the last 12 hours\n")
}

var duration string

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

func isSameDay(date1, date2 time.Time) bool {
	year1, month1, day1 := date1.Date()
	year2, month2, day2 := date2.Date()

	return year1 == year2 && month1 == month2 && day1 == day2
}

func isWithin(date1 time.Time, date2 time.Time, dur time.Duration) bool {
	return date1.Add(dur).After(date2)
}
