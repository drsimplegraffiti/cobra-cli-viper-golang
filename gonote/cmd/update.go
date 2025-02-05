package cmd

import (
	"database/sql"
	"fmt"
	"gonote/internal/store"
	"gonote/utils"
	"log"

	"github.com/spf13/cobra"
)


var noteID int
var newContent string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update note content",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()

		// Convert newContent to sql.NullString
		content := sql.NullString{
			String: newContent,
			Valid:  newContent != "", // If newContent is empty, set Valid to false
		}

		// Create the parameter object for UpdateNoteContent
		params := store.UpdateNoteContentParams{
			ID:      int32(noteID),
			Content: content,
		}

		// Call UpdateNoteContent with the correct parameter type
		err := queries.UpdateNoteContent(cmd.Context(), params)
		if err != nil {
			log.Fatalf("Failed to update note: %v", err)
		}
		fmt.Println("Note updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&noteID, "id", "i", 0, "ID of the note")
	updateCmd.Flags().StringVarP(&newContent, "content", "c", "", "New content")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("content")
}