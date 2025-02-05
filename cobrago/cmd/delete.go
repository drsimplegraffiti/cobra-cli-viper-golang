package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "strconv"
)

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete a task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please provide a task number to delete")
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

        tasks = append(tasks[:index-1], tasks[index:]...)
        viper.Set("tasks", tasks)
        viper.WriteConfig()
        fmt.Println("Task deleted.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
