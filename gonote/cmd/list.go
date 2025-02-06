// package cmd

// import (
// 	"fmt"
// 	"gonote/utils"
// 	"log"

// 	"github.com/spf13/cobra"
// )

// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "List all notes",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		queries := utils.GetDB()
// 		notes, err := queries.ListNotes(cmd.Context())
// 		if err != nil {
// 			log.Fatalf("Failed to fetch notes: %v", err)
// 		}
// 		for _, note := range notes {
// 			fmt.Printf("%d: [%s] %s\n", note.ID, note.Category.String, note.Title)
// 		}
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(listCmd)
// }
package cmd

import (
	"fmt"
	"gonote/utils"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
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

		// Prepare table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Category", "Title"})
		table.SetBorder(false)
		table.SetRowLine(true)

		// Set header colors
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		)

		// Set column colors
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor}, // ID
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},   // Category
			tablewriter.Colors{tablewriter.FgHiMagentaColor},                  // Title
		)

		// Populate table rows
		for _, note := range notes {
			table.Append([]string{
				fmt.Sprintf("%d", note.ID),
				note.Category.String,
				note.Title,
			})
		}

		// Set footer
		table.SetFooter([]string{"", "Total Notes", fmt.Sprintf("%d", len(notes))})
		table.SetFooterColor(
			tablewriter.Colors{},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		)

		// Render the table
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
