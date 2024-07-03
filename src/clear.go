package src

import (
	"os"
	"os/exec"
)

func Clear() {
	//time.Sleep(300 * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
