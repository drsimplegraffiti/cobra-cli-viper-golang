package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "strconv"
)

var doneCmd = &cobra.Command{
    Use:   "done",
    Short: "Mark a task as done",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please provide a task number to mark as done")
            return
        }

        index, err := strconv.Atoi(args[0])
        if err != nil || index < 1 {
            fmt.Println("Invalid task number")
            return
        }

        tasks := viper.GetStringSlice("tasks")
        if index > len(tasks) {
            fmt.Println("Task number out of range")
            return
        }

        task := tasks[index-1]
        doneTasks := viper.GetStringSlice("done_tasks")
        doneTasks = append(doneTasks, task)

        // Remove task from the list
        tasks = append(tasks[:index-1], tasks[index:]...)
        viper.Set("tasks", tasks)
        viper.Set("done_tasks", doneTasks)
        viper.WriteConfig()

        fmt.Println("Task marked as done:", task)
    },
}

func init() {
    rootCmd.AddCommand(doneCmd)
}
