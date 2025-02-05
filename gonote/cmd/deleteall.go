package cmd

import (
	"bufio"
	"fmt"
	"gonote/utils"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// deleteallCmd represents the deleteall command
var deleteallCmd = &cobra.Command{
	Use:   "deleteall",
	Short: "Delete all records",
	Long:  `This command deletes all categories and notes from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// Prompt user for confirmation
		fmt.Print("Are you sure you want to delete all records? This action cannot be undone (y/n): ")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)

		// Check if user confirmed with 'y' or 'yes'
		if confirm != "y" && confirm != "yes" {
			fmt.Println("Aborted: No records were deleted.")
			return
		}

		// Proceed with deletion
		queries := utils.GetDB()

		// Delete all notes
		if err := queries.DeleteAllNotes(cmd.Context()); err != nil {
			log.Fatalf("Failed to delete all notes: %v", err)
		}

		// Delete all categories
		if err := queries.DeleteAllCategories(cmd.Context()); err != nil {
			log.Fatalf("Failed to delete all categories: %v", err)
		}

		fmt.Println("✅ All records (notes and categories) have been deleted.")
	},
}

func init() {
	rootCmd.AddCommand(deleteallCmd)
}
