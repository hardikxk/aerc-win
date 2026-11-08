//go:build windows

package term

import (
	"os"
	"os/exec"

	"go.rockorager.dev/vaxis/widgets/term/pty"
)

func (vt *Model) startPty(cmd *exec.Cmd, width, height int) (*os.File, error) {
	return nil, pty.ErrUnsupported
}

