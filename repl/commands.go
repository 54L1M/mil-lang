package repl

import (
	"os"
	"os/exec"
)

func clearOutput() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
