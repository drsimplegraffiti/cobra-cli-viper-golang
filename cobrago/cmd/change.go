package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "strconv"
    "strings"
)

// updateCmd represents the "update" command
var updateCmd = &cobra.Command{
    Use:   "change", // Change this to "update" instead of "upsert"
    Short: "Update an existing task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            fmt.Println("Usage: todo update <task number> <new task description>")
            return
        }

        taskNum, err := strconv.Atoi(args[0])
        if err != nil || taskNum < 1 {
            fmt.Println("Invalid task number")
            return
        }

        newTask := strings.TrimSpace(args[1])
        tasks := viper.GetStringSlice("tasks")

        // Check if the task number is valid
        if taskNum > len(tasks) {
            fmt.Println("Error: Task number out of range")
            return
        }

        // Update the task
        tasks[taskNum-1] = newTask
        viper.Set("tasks", tasks)

        if err := viper.WriteConfig(); err != nil {
            fmt.Println("Error writing to config:", err)
            return
        }

        fmt.Printf("Task %d updated: %s\n", taskNum, newTask)
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
}
