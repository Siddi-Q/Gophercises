package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/db"
	"github.com/fatih/color"
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
		now := time.Now()
		tasks, err := db.ReadSomeTasks(true)

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		duration = strings.ToLower(duration)
		switch duration {
		case "forever":
			completedCmdOutput("You have not completed any tasks!", "You have completed the following tasks:", tasks)
		case "today":
			var filteredTasks []db.Task
			for _, task := range tasks {
				if isSameDay(task.CompletedDate, now) {
					filteredTasks = append(filteredTasks, task)
				}
			}

			completedCmdOutput("You have not completed any tasks today!", "You have completed the following tasks today:", filteredTasks)
		case "24h":
			var filteredTasks []db.Task
			for _, task := range tasks {
				if isWithin(task.CompletedDate, now, time.Hour*24) {
					filteredTasks = append(filteredTasks, task)
				}
			}

			completedCmdOutput("You have not completed any tasks in the last 24 hours!", "You have completed the following tasks in the last 24 hours:", filteredTasks)
		case "12h":
			var filteredTasks []db.Task
			for _, task := range tasks {
				if isWithin(task.CompletedDate, now, time.Hour*12) {
					filteredTasks = append(filteredTasks, task)
				}
			}

			completedCmdOutput("You have not completed any tasks in the last 12 hours!", "You have completed the following tasks in the last 12 hours:", filteredTasks)
		default:
			color.New(color.FgRed).Println("Unrecognized duration. Please try again!")
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

func outputTasks(tasks []db.Task) {
	color.Set(color.FgGreen)
	for i, task := range tasks {
		fmt.Printf("%d. %s | Completed on %v\n", i+1, task.Description, task.CompletedDate.Format("Monday, January 2, 2006 03:04:05 PM"))
	}
	color.Unset()
}

func completedCmdOutput(noTasksOuputString string, tasksOutputString string, tasks []db.Task) {
	if len(tasks) == 0 {
		color.New(color.FgYellow).Println(noTasksOuputString)
	} else {
		fmt.Println(tasksOutputString)
		outputTasks(tasks)
	}
}
