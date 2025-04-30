package cmd

import (
	"archive/zip"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var backupSqlCmd = &cobra.Command{
	Use:   "backupsql",
	Short: "Realiza o backup de um banco de dados SQL Server",
	Run:   executarBackupSQL,
}

func init() {
	rootCmd.AddCommand(backupSqlCmd)
}

func executarBackupSQL(cmd *cobra.Command, args []string) {
	var senha string
	var porta string

	promptSenha := &survey.Password{
		Message: "Senha do banco de dados:",
	}
	survey.AskOne(promptSenha, &senha)

	promptPorta := &survey.Input{
		Message: "Porta do banco (padrão 5433):",
		Default: "5433",
	}
	survey.AskOne(promptPorta, &porta)

	host := "localhost"
	user := "sa"

	// Conexão com master para listar bancos
	masterConn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=master", user, senha, host, porta)
	db, err := sql.Open("sqlserver", masterConn)
	if err != nil {
		color.Red("Erro ao conectar ao banco master: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM sys.databases WHERE name NOT IN ('master', 'tempdb', 'model', 'msdb')")
	if err != nil {
		color.Red("Erro ao listar bancos: %v", err)
		os.Exit(1)
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			databases = append(databases, name)
		}
	}

	if len(databases) == 0 {
		color.Red("Nenhum banco disponível para backup foi encontrado.")
		os.Exit(1)
	}

	var dbSelecionado string
	promptDB := &survey.Select{
		Message: "Escolha o banco de dados:",
		Options: databases,
	}
	survey.AskOne(promptDB, &dbSelecionado)

	realizarBackup(dbSelecionado, senha, host, porta, user)
}

func realizarBackup(dbName, senha, host, porta, user string) {
	now := time.Now()
	backupDir := filepath.Join("C:/Softcom/backup", now.Format("2006"), now.Format("01"), now.Format("02"))
	os.MkdirAll(backupDir, os.ModePerm)

	// Caminho TEMP acessível ao SQL Server
	bakPath := filepath.Join("C:/Softcom/temp", fmt.Sprintf("%s.bak", dbName))
	os.MkdirAll("C:/Softcom/temp", os.ModePerm)

	zipPath := filepath.Join(backupDir, fmt.Sprintf("%s.zip", dbName))
	query := fmt.Sprintf("BACKUP DATABASE [%s] TO DISK = N'%s' WITH INIT", dbName, bakPath)
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, senha, host, porta, dbName)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		color.Red("Erro ao conectar ao banco: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		color.Red("Erro ao realizar o backup do banco '%s': %v", dbName, err)
		os.Exit(1)
	}

	color.Green("Backup do banco '%s' realizado com sucesso!", dbName)

	err = compactarZip(bakPath, zipPath)
	if err != nil {
		color.Red("Erro ao compactar o backup: %v", err)
		os.Exit(1)
	}

	os.Remove(bakPath)
	color.Green("Backup compactado com sucesso!")
	color.Cyan("Arquivo salvo em: %s", zipPath)
}

func compactarZip(origem, destino string) error {
	zipFile, err := os.Create(destino)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	w := zip.NewWriter(zipFile)
	defer w.Close()

	f, err := os.Open(origem)
	if err != nil {
		return err
	}
	defer f.Close()

	info, _ := f.Stat()
	header, _ := zip.FileInfoHeader(info)
	header.Name = filepath.Base(origem)
	header.Method = zip.Deflate

	writer, err := w.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, f)
	return err
}
