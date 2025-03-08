package cmd

import (
	"fmt"
	"log"

	"github.com/DevAthhh/task-tracker-cli/app/task"
	"github.com/spf13/cobra"
)

// TODO:
// 	- remove dependencies

func newAddCommand() *cobra.Command {
	var taskAdd = &cobra.Command{
		Use:   "add",
		Short: "This command adds a task",
		Long:  "This command adds a task with user descriptions to JSON file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			desc := args[0]
			res, err := task.TaskAddFunc(desc)
			if err != nil {
				log.Fatalf("error adding task: %v", err)
			}
			fmt.Println(res)
		},
	}
	return taskAdd
}
func newUpdateCommand() *cobra.Command {
	var taskUpdate = &cobra.Command{
		Use:   "update",
		Short: "This command update a desc of task",
		Long:  "This command update a desc of task",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := task.TaskUpdateFunc(args[0], args[1]); err != nil {
				log.Fatalf("error updating task: %v", err)
			}
		},
	}
	return taskUpdate
}
func newDeleteCommand() *cobra.Command {
	var taskDelete = &cobra.Command{
		Use:   "delete",
		Short: "This command delete a task",
		Long:  "This command delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := task.TaskDeleteFunc(args[0]); err != nil {
				log.Fatalf("error deleting task: %v", err)
			}
		},
	}
	return taskDelete
}

func newMarkInProgressCommand() *cobra.Command {
	var taskMarkInProgress = &cobra.Command{
		Use:   "mark-in-progress",
		Short: "This command marks the task as 'in progress'",
		Long:  "This command marks the task as 'in progress'",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := task.TaskMarkStatusFunc(args[0], "in progress"); err != nil {
				log.Fatalf("error marking task: %v", err)
			}
		},
	}
	return taskMarkInProgress
}

func newMarkDoneCommand() *cobra.Command {
	var taskMarkDone = &cobra.Command{
		Use:   "mark-done",
		Short: "This command marks the task as 'done'",
		Long:  "This command marks the task as 'done'",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := task.TaskMarkStatusFunc(args[0], "done"); err != nil {
				log.Fatalf("error marking task: %v", err)
			}
		},
	}
	return taskMarkDone
}

func newListCommand() *cobra.Command {
	var taskList = &cobra.Command{
		Use:   "list",
		Short: "this command returns a list of all tasks.",
		Long:  "this command returns a list of all tasks.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var cond string
			if len(args) < 1 {
				cond = ""
			} else {
				cond = args[0]
			}
			tasks, err := task.TaskListFunc(cond)
			if err != nil {
				log.Fatalf("error list task: %v", err)
			}
			for _, task := range tasks {
				fmt.Println("-------------------------------------------------")
				fmt.Println("ID:", task.ID)
				fmt.Println("Desc:", task.Desc)
				fmt.Println("Status:", task.Status)
				fmt.Println("CreatedAt:", task.CreatedAt)
				fmt.Println("UpdatedAt:", task.UpdatedAt)
			}
		},
	}
	return taskList
}
