package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Gerencia os logs do sistema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Comando de gerenciamento de logs executado.")
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
