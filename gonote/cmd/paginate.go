package cmd

import (
	"fmt"
	"gonote/internal/store"
	"gonote/utils"
	"log"
	"github.com/spf13/cobra"
)

var categoryFilter string
var titleFilter string
var page int
var pageSize int

var paginateCmd = &cobra.Command{
	Use:   "paginate",
	Short: "List notes with pagination",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()

		// Set default page size if not provided
		if pageSize == 0 {
			pageSize = 10
		}

		// Calculate offset for pagination
		offset := (page - 1) * pageSize

		// Build filter parameters with the correct types (cast to int32)
		params := store.PaginateNotesParams{
			Column1: categoryFilter,
			Column2:    titleFilter,
			Limit:    int32(pageSize),  // Cast to int32
			Offset:   int32(offset),    // Cast to int32
		}

		// Fetch notes with pagination and filters
		notes, err := queries.PaginateNotes(cmd.Context(), params)
		if err != nil {
			log.Fatalf("Failed to fetch notes: %v", err)
		}
		for _, note := range notes {
			fmt.Printf("%d: [%s] %s (Created at: %s)\n", note.ID, note.Category, note.Title, note.CreatedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(paginateCmd)

	// Add flags for filtering notes and pagination
	paginateCmd.Flags().StringVarP(&categoryFilter, "category", "c", "", "Filter notes by category")
	paginateCmd.Flags().StringVarP(&titleFilter, "title", "t", "", "Filter notes by title")
	paginateCmd.Flags().IntVarP(&page, "page", "p", 1, "Page number")
	paginateCmd.Flags().IntVarP(&pageSize, "pageSize", "s", 10, "Number of notes per page")
}
