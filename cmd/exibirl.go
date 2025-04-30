package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var exibirLCmd = &cobra.Command{
	Use:   "exibirl",
	Short: "Exibe os logs recentes da aplicação",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exibindo logs recentes...")
		// Lógica para ler e exibir logs
	},
}

func init() {
	rootCmd.AddCommand(exibirLCmd)
}
