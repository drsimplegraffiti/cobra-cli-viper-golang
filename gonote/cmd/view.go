package cmd

import (
	"fmt"
	"gonote/utils"
	"log"

	"github.com/spf13/cobra"
)

var viewID int

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View note details",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()
		note, err := queries.GetNote(cmd.Context(), int32(viewID))
		if err != nil {
			log.Fatalf("Failed to fetch note: %v", err)
		}
		fmt.Printf("Title: %s\nCategory: %s\nContent:\n%s\n", note.Title, note.Category.String, note.Content)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().IntVarP(&viewID, "id", "i", 0, "ID of the note")
	viewCmd.MarkFlagRequired("id")
}
