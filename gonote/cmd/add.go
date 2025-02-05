// package cmd

// import (
// 	"database/sql"
// 	"fmt"
// 	"gonote/internal/store"
// 	"gonote/utils"
// 	"log"

// 	"github.com/spf13/cobra"
// )

// var title, content string
// var categoryID int

// var addCmd = &cobra.Command{
// 	Use:   "add",
// 	Short: "Add a new note",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		queries := utils.GetDB()
// 		note, err := queries.CreateNote(cmd.Context(), store.CreateNoteParams{
// 			Title:      title,
// 			CategoryID: sql.NullInt32{Int32: int32(categoryID), Valid: true},
// 			Content:    sql.NullString{String: content, Valid: content != ""},
// 		})
// 		if err != nil {
// 			log.Fatalf("Failed to add note: %v", err)
// 		}
// 		fmt.Printf("Added Note: %s (ID: %d)\n", note.Title, note.ID)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(addCmd)
// 	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the note")
// 	addCmd.Flags().StringVarP(&content, "content", "x", "", "Title of the note")
// 	addCmd.Flags().IntVarP(&categoryID, "category", "c", 0, "Category ID")
// 	addCmd.MarkFlagRequired("title")
// }
package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"gonote/internal/store"
	"gonote/utils"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// // Prompt for title
		// fmt.Print("Enter note title: ")
		// title, _ := reader.ReadString('\n')
		// title = strings.TrimSpace(title)

		// if title == "" {
		// 	fmt.Println("Title cannot be empty!")
		// 	return
		// }

		var title string
		for {
			fmt.Print("Enter note title: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)

			if title != "" {
				break // Exit loop if valid title is provided
			}

			fmt.Println("❌ Title cannot be empty! Please enter a valid title.")
		}

		// Prompt for content (multi-line input)
		fmt.Println("Enter note content (press ENTER twice to finish):")
		var contentLines []string
		for {
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if line == "" { // Stop input on empty line
				break
			}

			contentLines = append(contentLines, line)
		}
		content := strings.Join(contentLines, "\n")

		// Prompt for category
		fmt.Print("Enter category ID (or press ENTER to skip): ")
		categoryInput, _ := reader.ReadString('\n')
		categoryInput = strings.TrimSpace(categoryInput)

		var categoryID sql.NullInt32
		if categoryInput != "" {
			if num, err := strconv.Atoi(categoryInput); err == nil {
				categoryID.Int32 = int32(num)
				categoryID.Valid = true
			} else {
				fmt.Println("Invalid category ID. Using default (NULL).")
			}
		}

		queries := utils.GetDB()

		// Validate category ID if provided
		if categoryID.Valid {
			_, err := queries.GetCategoryById(cmd.Context(), categoryID.Int32)
			if err != nil {
				fmt.Println("❌ Invalid category ID! Please enter a valid category.")
				return
			}
		}


		// Save note
		// queries := utils.GetDB()
		note, err := queries.CreateNote(cmd.Context(), store.CreateNoteParams{
			Title:      title,
			CategoryID: categoryID,
			Content:    sql.NullString{String: content, Valid: content != ""},
		})
		if err != nil {
			log.Fatalf("Failed to add note: %v", err)
		}

		fmt.Printf("\n✅ Note added: %s (ID: %d)\n", note.Title, note.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
