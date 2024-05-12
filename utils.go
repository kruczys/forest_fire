package main

import (
	"os"
	"os/exec"
)

func clearScreen() { // funkcja pomocnicza czyszczaca terminal
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
