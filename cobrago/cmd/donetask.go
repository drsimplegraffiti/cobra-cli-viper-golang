package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var donetasks =  &cobra.Command{
	Use:   "dones",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        doneTasks := viper.GetStringSlice("done_tasks")
        if len(doneTasks) == 0 {
            fmt.Println("No tasks found.")
            return
        }
        fmt.Println("Your done tasks:")
        for i, doneTasks := range doneTasks {
            fmt.Printf("%d. %s\n", i+1, doneTasks)
        }
    },
}


func init() {
    rootCmd.AddCommand(donetasks)
}