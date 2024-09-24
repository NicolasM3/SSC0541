package ui

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	default:
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
