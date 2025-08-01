/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Abiji-2020/PesudoCLI/internal/config"
	"github.com/Abiji-2020/PesudoCLI/internal/gemini"
	"github.com/Abiji-2020/PesudoCLI/internal/redisclient"
	"github.com/spf13/cobra"
)

var (
	RedisClient  *redisclient.RedisClient
	GeminiClient *gemini.GeminiClient
)

var rootCmd = &cobra.Command{
	Use:   "PesudoCLI",
	Short: "Simple man page AI helper",
	Long: `PesudoCLI is a simple command line interface that helps you with
a simple response for the commands you needed in your terminao. It comes with 
the help from Redis to store the commands and responses, and uses Gemini to 
make the response more accurate from the context of the question and the vector store 
`,
	Run: func(cmd *cobra.Command, args []string) {},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔧 [SETUP] Loading config and connecting to Redis...")
		config, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		RedisClient = redisclient.NewRedisClient(config.RedisAddr)
		GeminiClient, err = gemini.NewGeminiClient(config.GEMINI_API_KEY)
		if err != nil {
			log.Fatalf("Error initializing Gemini client: %v", err)
		}
		fmt.Println("✅ [SETUP] Redis and Gemini clients initialized successfully")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
