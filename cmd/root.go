package cmd

import (
	"fmt"
	"os"

	"github.com/RobertoFilho0100/softcli/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "softcli",
	Short: "SoftCLI - Ferramenta oficial da Softcom",
	Long:  `SoftCLI é uma ferramenta de linha de comando para facilitar o dia a dia do time de suporte da Softcom.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.PrintHeader()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
