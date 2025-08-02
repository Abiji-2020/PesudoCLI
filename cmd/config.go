/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"github.com/Abiji-2020/PesudoCLI/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Command to manage configuration settings",
	Long: `
The config command allows you to manage the coniguration
settings for PesudoCLI. You can view or modify settings using this command. 
Currently, it has the configuration of the following: 

	- Redis address for the Redis client
	- Gemini API key for the Gemini client
	- Gemini embedding model for the Gemini client
	- Gemini chat model for the Gemini client`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.SaveConfig(&config.Config{
			RedisAddr:              cmd.Flag("redis-addr").Value.String(),
			GEMINI_API_KEY:         cmd.Flag("gemini-api-key").Value.String(),
			GEMINI_EMBEDDING_MODEL: cmd.Flag("gemini-embedding-model").Value.String(),
			GEMINI_CHAT_MODEL:      cmd.Flag("gemini-chat-model").Value.String(),
			IndexName:              cmd.Flag("index-name").Value.String(),
		})
		if err != nil {
			cmd.Println("Error saving configuration:", err)
			return
		}
		cmd.Println("Configuration saved successfully.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("redis-addr", "r", "localhost:6379", "Redis address (default: localhost:6379)")
	configCmd.Flags().StringP("gemini-api-key", "g", "", "Gemini API key")
	configCmd.Flags().StringP("gemini-embedding-model", "e", "gemini-embedding-001", "Gemini embedding model (default: gemini-embedding-001)")
	configCmd.Flags().StringP("gemini-chat-model", "c", "gemini-2.0-flash", "Gemini chat model (default: gemini-2.0-flash)")
	configCmd.Flags().StringP("index-name", "i", "pesudo_index", "Index name for the vector store (default: pesudo_index)")
	configCmd.MarkFlagRequired("gemini-api-key")
}
