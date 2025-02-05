package cmd

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "fmt"
    "os"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Short: "A simple CLI TODO app",
	Long: "A simple command line TODO app built with Cobra and Viper",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    viper.SetConfigName("todo")
    viper.SetConfigType("json")
    viper.AddConfigPath("./data")

    if err := viper.ReadInConfig(); err != nil {
        viper.SafeWriteConfig()
    }
}