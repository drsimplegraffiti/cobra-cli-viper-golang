package cmd

import (
	"fmt"
	"gonote/utils"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()
		notes, err := queries.ListNotes(cmd.Context())
		if err != nil {
			log.Fatalf("Failed to fetch notes: %v", err)
		}
		for _, note := range notes {
			fmt.Printf("%d: [%s] %s\n", note.ID, note.Category.String, note.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
