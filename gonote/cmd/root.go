package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonote",
	Short: "Go note",
	Long:  `Go note for me`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// initConfig loads configuration from file
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Set to the directory where config.yaml is located

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// fmt.Println("Using config:", viper.ConfigFileUsed())
}

func init() {
	// Load config before executing any commands
	cobra.OnInitialize(initConfig)

	// Define persistent flags for global use
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gonote.yaml)")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	rootCmd.PersistentFlags().StringP("author", "a", "Abayomi Ogunnusi", "Author name for copyright attribution")
	viper.SetDefault("author", "Abayomi Ogunnusi <abayomiogunnusi@gmail.com>")
	viper.SetDefault("license", "apache")

	// Define local flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
