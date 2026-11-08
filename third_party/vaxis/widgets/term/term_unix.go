//go:build !windows

package term

import (
	"os"
	"os/exec"
	"syscall"

	"go.rockorager.dev/vaxis/widgets/term/pty"
)

func (vt *Model) startPty(cmd *exec.Cmd, width, height int) (*os.File, error) {
	winsize := pty.Winsize{
		Cols: uint16(width),
		Rows: uint16(height),
	}
	return pty.StartWithAttrs(
		cmd,
		&winsize,
		&syscall.SysProcAttr{
			Setsid:  true,
			Setctty: true,
			Ctty:    1,
		},
	)
}

