package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrintHeader() {
	logo := `
   ███████╗ ██████╗ ███████╗████████╗ ██████╗██╗     ██╗
   ██╔════╝██╔═══██╗██╔════╝╚══██╔══╝██╔═══██╗██║     ██║
   █████╗  ██║   ██║███████╗   ██║   ██║   ██║██║     ██║
   ██╔══╝  ██║   ██║╚════██║   ██║   ██║   ██║██║     ██║
   ███████╗╚██████╔╝███████║   ██║   ╚██████╔╝███████╗███████╗
   ╚══════╝ ╚═════╝ ╚══════╝   ╚═╝    ╚═════╝ ╚══════╝╚══════╝

                      SoftCLI - Ferramenta oficial da Softcom
`
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00BCD4")).
		Bold(true).
		PaddingTop(1).
		PaddingBottom(1)

	fmt.Println(style.Render(logo))
}
