package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks := viper.GetStringSlice("tasks")
        if len(tasks) == 0 {
            fmt.Println("No tasks found.")
            return
        }
        fmt.Println("Your tasks:")
        for i, task := range tasks {
            fmt.Printf("%d. %s\n", i+1, task)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
