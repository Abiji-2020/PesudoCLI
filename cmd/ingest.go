/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"github.com/Abiji-2020/PesudoCLI/internal/config"
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
		cmd.Println("🔧 [INGEST] Ingesting data and embedding it into the vector store")
		embeddedValue, err := GeminiClient.Embed("Hello, this is a test embedding.", config.GEMINI_EMBEDDING_MODEL)
		if err != nil {
			cmd.Println("Error embedding data:", err)
			return
		}
		err = RedisClient.AddDocument("test_id", "test_command", "test_os", "Hello, this is a test document.", embeddedValue)
		if err != nil {
			cmd.Println("Error adding document to Redis:", err)
			return
		}
		cmd.Println("✅ [INGEST] Data ingested and embedded successfully")

	},
}

func init() {
	rootCmd.AddCommand(ingestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
