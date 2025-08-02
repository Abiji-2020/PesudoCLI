/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"github.com/Abiji-2020/PesudoCLI/internal/config"
	"github.com/Abiji-2020/PesudoCLI/pkg/io"
	"github.com/spf13/cobra"
)

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest data and embed it into the vector store",
	Long: `The ingest command allows you to ingest data and embed it into the vector store.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			cmd.Println("Error loading configuration:", err)
			return
		}
		// Load Documents
		docs, err := io.LoadComandDocs()
		if err != nil {
			cmd.Println("Error loading command documents:", err)
			return
		}

		cmd.Println("🔧 [INGEST] Ingesting data and embedding it into the vector store")
		indexName, _ := cmd.Flags().GetString("index")
		if indexName == "" {
			cmd.Println("Please provide an index name using --index flag")
			return
		}
		if err := RedisClient.CreateVectorIndex(indexName, 3072); err != nil {
			cmd.Println("Error creating vector index:", err)
			return
		}
		cmd.Println("🔧 [INGEST] Vector index created successfully")
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			cmd.Println("Error retrieving limit value:", err)
			return
		}
		if limit <= 0 {
			cmd.Println("Please provide a valid limit using --limit flag")
			return
		}
		if len(docs) > limit {
			docs = docs[:limit] // Limit the number of documents to ingest
		}
		if len(docs) > 800 {
			docs = docs[:800] // Limit to 800 to avoid limiting issues
		}
		embeddedValue, err := GeminiClient.Embed(docs, config.GEMINI_EMBEDDING_MODEL)
		if err != nil {
			cmd.Println("Error embedding data:", err)
			return
		}
		cmd.Println("🔧 [INGEST] Data embedded successfully")
		err = RedisClient.AddDocument(embeddedValue)
		if err != nil {
			cmd.Println("Error adding document to Redis:", err)
			return
		}
		cmd.Println("✅ [INGEST] Data ingested and embedded successfully")

	},
}

func init() {
	rootCmd.AddCommand(ingestCmd)

	ingestCmd.Flags().StringP("index", "i", "pesudo_index", "Index name for the vector store (default: pesudo_index)")
	ingestCmd.Flags().IntP("limit", "l", 800, "Limit the number of documents to ingest (default: 800)")
}
