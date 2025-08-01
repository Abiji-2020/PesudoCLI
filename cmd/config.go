/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Command to manage configuration settings",
	Long: `
The config command allows you to manage the coniguration
settings for PesudoCLI. You can view or modify settings using this command. 
Currently, it has the configuration of the following: 

	- Redis address for the Redis client`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("config called")

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
