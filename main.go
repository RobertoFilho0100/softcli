package main

import (
	"fmt"
	"os"

	"github.com/RobertoFilho0100/softcli/cmd"
	"github.com/RobertoFilho0100/softcli/internal/ui"
)

func main() {
	ui.PrintHeader()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
