package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from a URL",
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url") // Get URL from Viper config
		if url == "" {
			fmt.Println("Error: URL not provided. Use --url or set it in the config file.")
			return
		}

		// Make an HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching URL: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}

		

		fmt.Printf("Response from %s:\n%s\n", url, string(body))
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
