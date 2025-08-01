/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A command to create the Index and vector store",
	Long: `The create command allows you to create a new index and vector store in Redis.
You can specify the index name and dimensions for the vector store.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔧 [SETUP] Creating Vector Index")
		indexName, _ := cmd.Flags().GetString("index")
		dim, _ := cmd.Flags().GetInt("dim")
		if indexName == "" {
			fmt.Println("Please provide an index name using --index flag")
			return
		}
		if dim <= 0 {
			fmt.Println("Please provide a valid dimension using --dim flag")
			return
		}
		if err := RedisClient.CreateVectorIndex(indexName, dim); err != nil {
			fmt.Printf("Error creating vector index: %v\n", err)
			return
		}
		fmt.Printf("Vector index '%s' created with dimension %d\n", indexName, dim)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("index", "i", "", "Name of the index to create")
	initCmd.Flags().IntP("dim", "d", 3072, "Dimensions for the vector store")
	initCmd.MarkFlagRequired("index")
}
