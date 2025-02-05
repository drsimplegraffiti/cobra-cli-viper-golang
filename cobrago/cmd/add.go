package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "strings"
)

// addCmd represents the "add" command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please provide a task description")
            return
        }

        task := strings.TrimSpace(args[0])
        tasks := viper.GetStringSlice("tasks")

        // Check if the task already exists
        for _, t := range tasks {
            if strings.EqualFold(t, task) { // Case-insensitive comparison
                fmt.Println("Error: Task already exists!")
                return
            }
        }

        // Add the new task
        tasks = append(tasks, task)
        viper.Set("tasks", tasks)
        if err := viper.WriteConfig(); err != nil {
            fmt.Println("Error writing to config:", err)
            return
        }

        fmt.Println("Task added:", task)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
