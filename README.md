
# SoftCLI

**SoftCLI** é uma ferramenta de linha de comando (CLI) desenvolvida para realizar tarefas de backup automatizado de bancos de dados SQL Server e sistemas locais (como o front-end em Access), com recursos de compactação e organização dos arquivos gerados. O projeto foi desenvolvido para facilitar o processo de backup e permitir a automação em ambientes corporativos.

## Índice

- [Visão Geral](#visão-geral)
- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Comandos](#comandos)
  - [backup](#comando-backup)
- [Estrutura do Projeto](#estrutura-do-projeto)

---

## Visão Geral

O **SoftCLI** oferece um comando simples para realizar o backup do banco de dados SQL Server e do sistema de front-end (usando Access), compactando os arquivos gerados em um arquivo `.zip`. O processo é totalmente automatizado, incluindo:

1. **Conexão com o banco de dados**: O comando `backup` se conecta ao banco de dados SQL Server e solicita as credenciais de acesso, como senha e porta.
2. **Backup do banco de dados**: O banco selecionado é copiado para um arquivo `.bak`.
3. **Backup do front-end**: O arquivo de banco de dados Access é copiado.
4. **Compactação dos arquivos**: Todos os arquivos gerados são compactados em um arquivo `.zip` para facilitar o armazenamento e envio.
5. **Organização dos arquivos**: Os backups são organizados por data e salvos em diretórios específicos.

---

## Requisitos

- **Go 1.18+**: Para compilar o projeto.
- **SQL Server**: A ferramenta foi projetada para trabalhar com o banco de dados SQL Server.
- **Access Database**: O backup do front-end utiliza o arquivo `.accdb` do Access.
- **Sistema operacional**: A CLI foi testada principalmente em Windows, mas deve funcionar em outros sistemas com a devida adaptação de caminhos e permissões.

---

## Instalação

Para começar a usar o SoftCLI, siga os passos abaixo:

### 1. Clone o repositório

Clone o repositório para o seu ambiente local:

```bash
git clone https://github.com/RobertoFilho0100/softcli.git
cd softcli
```

### 2. Instale as dependências

O projeto usa o **Go Modules** para gerenciamento de dependências, então basta rodar:

```bash
go mod tidy
```

### 3. Compile o projeto

Para compilar a aplicação, execute o seguinte comando:

```bash
go build -o softcli
```

Isso criará o executável `softcli` na raiz do projeto.

### 4. (Opcional) Instalar globalmente

Se quiser instalar o SoftCLI globalmente no seu sistema, mova o arquivo compilado para um diretório que esteja no `PATH`, como `/usr/local/bin` no Linux/Mac ou `C:\Windows\System32` no Windows.

---

## Comandos

### `backupsql`

Este comando realiza o backup do banco de dados SQL Server e do sistema de front-end, compactando tudo em um arquivo `.zip`.

#### Uso:

```bash
softcli backupsql
```

#### Descrição:

- O comando solicitará a **senha** do banco de dados SQL Server.
- Em seguida, pedirá a **porta** (padrão: `5433`).
- O usuário poderá selecionar o **banco de dados** a ser copiado.
- O banco selecionado será copiado para um arquivo `.bak`.
- O arquivo de front-end Access (`Softshop.accdb`) será copiado.
- Todos os arquivos serão compactados em um arquivo `.zip` com data e hora no nome.

#### Exemplo de saída:

```
   _____        __ _    _____ _      _____ 
  / ____|      / _| |  / ____| |    |_   _|
 | (___   ___ | |_| |_| |    | |      | |  
  \___ \ / _ \|  _| __| |    | |      | |  
  ____) | (_) | | | |_| |____| |____ _| |_ 
 |_____/ \___/|_|  \__|\_____|______|_____|
        SoftLogger - Backup SQL

Senha do banco de dados: ********
Porta do banco (padrão 5433): 5433
Selecionar o banco de dados: 
  - Banco1
  - Banco2
  - Banco3

Backup realizado com sucesso!
Arquivo gerado: C:/Softcom/backup/2025/04/30/backup_Banco1.zip
```

---

## Estrutura do Projeto

O projeto possui a seguinte estrutura de diretórios:

```
softcli/
├── cmd/                 # Comandos da CLI (backup, etc.)
│   ├── backup.go        # Implementação do comando "backup"
│   └── root.go          # Comando raiz da CLI
│   └── backupfront.go   # Comando raiz da CLI
│   └── backupsql.go     # Comando raiz da CLI
│   └── logs.go          # Comando raiz da CLI
├── internal/            # Código interno do projeto
│   └── ui/              # Funções auxiliares (ex: exibição de header)
├── go.mod               # Gerenciamento de dependências Go
├── go.sum               # Hashes das dependências
└── README.md            # Este arquivo
```
---

### Observações

- Certifique-se de que o SQL Server está rodando e acessível antes de tentar realizar o backup.
- A CLI exige permissões de leitura e escrita nos diretórios de destino e de onde o backup do front-end será copiado.

---
