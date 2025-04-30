package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var backupFrontCmd = &cobra.Command{
	Use:   "backupfront",
	Short: "Realiza backup do frontend da aplicação",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Backup do frontend iniciado...")
		// Implementar lógica do backup do frontend aqui
		fmt.Println("Backup do frontend concluído com sucesso.")
	},
}

func init() {
	rootCmd.AddCommand(backupFrontCmd)
}
