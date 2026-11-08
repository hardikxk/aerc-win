//go:build windows

package completer

import (
	"os/exec"
)

func setProcessGroup(cmd *exec.Cmd) {
}

func killProcessGroup(cmd *exec.Cmd) {
	if cmd != nil && cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
}

