package cmd

import (
	"fmt"
	"gonote/utils"
	"log"

	"github.com/spf13/cobra"
)

var categoryName string

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage categories",
}

var addCategoryCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new category",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()
		category, err := queries.CreateCategory(cmd.Context(), categoryName)
		if err != nil {
			log.Fatalf("Failed to create category: %v", err)
		}
		fmt.Printf("Created Category: %s (ID: %d)\n", category.Name, category.ID)
	},
}

var listCategoriesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all categories",
	Run: func(cmd *cobra.Command, args []string) {
		queries := utils.GetDB()
		categories, err := queries.ListCategories(cmd.Context())
		if err != nil {
			log.Fatalf("Failed to fetch categories: %v", err)
		}
		for _, category := range categories {
			fmt.Printf("%d: %s\n", category.ID, category.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.AddCommand(addCategoryCmd, listCategoriesCmd)

	addCategoryCmd.Flags().StringVarP(&categoryName, "name", "n", "", "Category name")
	addCategoryCmd.MarkFlagRequired("name")
}
