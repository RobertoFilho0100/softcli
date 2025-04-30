package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var backupSqlCmd = &cobra.Command{
	Use:   "backupsql",
	Short: "Realiza backup do banco SQL Server",
	Run: func(cmd *cobra.Command, args []string) {
		realizarBackup()
	},
}

func realizarBackup() {
	server := "localhost,5433" // ou 1433
	user := "sa"

	var senha string
	fmt.Print("Digite a senha do SQL Server: ")
	fmt.Scanln(&senha)

	// Conectar
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;encrypt=disable", server, user, senha)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Erro de conexão:", err)
		return
	}
	defer db.Close()

	// Listar bancos
	rows, err := db.Query("SELECT name FROM sys.databases WHERE database_id > 4")
	if err != nil {
		fmt.Println("Erro ao listar bancos:", err)
		return
	}
	defer rows.Close()

	var bancos []string
	for rows.Next() {
		var nome string
		rows.Scan(&nome)
		bancos = append(bancos, nome)
	}

	// Menu interativo
	prompt := promptui.Select{
		Label: "Selecione o banco para backup",
		Items: bancos,
		Size:  10,
	}
	_, bancoEscolhido, err := prompt.Run()
	if err != nil {
		fmt.Println("Erro ao selecionar banco:", err)
		return
	}

	// Spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Realizando backup..."
	s.Start()

	// Executar BACKUP
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	backupPath := fmt.Sprintf("C:\\Temp\\%s_%d.bak", bancoEscolhido, time.Now().Unix())
	_, err = db.ExecContext(ctx, fmt.Sprintf("BACKUP DATABASE [%s] TO DISK = '%s'", bancoEscolhido, backupPath))
	s.Stop()

	if err != nil {
		fmt.Println("Erro ao fazer backup:", err)
		return
	}

	fmt.Println("✅ Backup concluído em:", backupPath)
}
