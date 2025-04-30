package ui

import "github.com/fatih/color"

func PrintHeader() {
	color.Cyan(`
   _____        __ _    _____ _      _____ 
  / ____|      / _| |  / ____| |    |_   _|
 | (___   ___ | |_| |_| |    | |      | |  
  \___ \ / _ \|  _| __| |    | |      | |  
  ____) | (_) | | | |_| |____| |____ _| |_ 
 |_____/ \___/|_|  \__|\_____|______|_____|          
        SoftLogger - Backup SQL
	`)
}
