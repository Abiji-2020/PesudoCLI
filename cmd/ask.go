/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"github.com/Abiji-2020/PesudoCLI/internal/config"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "A command to ask questions",
	Long: `A simple command to ask questions and get answers.
This command is designed to interact with the system and provide answers based on the context.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]
		if question == "" {
			cmd.Println("Please provide a question to ask.")
			return
		}

		index := cmd.Flag("index").Value.String()
		topK := cmd.Flag("topk").Value.String()

		config, err := config.LoadConfig()
		if err != nil {
			cmd.Println("Error loading configuration:", err)
			return
		}
		query_vec, err := GeminiClient.EmbedQuestion(question, config.GEMINI_EMBEDDING_MODEL)
		if err != nil {
			cmd.Println("Error embedding question:", err)
			return
		}
		queryResult, err := RedisClient.QuerySearch(index, topK, query_vec)
		if err != nil {
			cmd.Println("Error querying Redis:", err)
			return
		}
		answer, err := GeminiClient.AskQuestion(question, config.GEMINI_CHAT_MODEL, queryResult)
		if err != nil {
			cmd.Println("Error asking question:", err)
			return
		}
		cmd.Println("-----------------------------------------")
		cmd.Println("OS:", answer.Os)
		cmd.Println("Command:", answer.Command)
		cmd.Println("Explanation:", answer.Explanation)
		cmd.Println("-----------------------------------------")
		cmd.Println("Detailed Explanation:\n", answer.Answer)
		cmd.Println("-----------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
	askCmd.Flags().StringP("index", "i", "pesudo_index", "Index name for the vector store (default: pesudo_index)")
	askCmd.Flags().IntP("topk", "k", 3, "Number of top results to return (default: 3)")
}
