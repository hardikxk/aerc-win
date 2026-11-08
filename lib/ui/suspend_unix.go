//go:build !windows

package ui

import (
	"os"
	"os/signal"
	"syscall"
)

func suspendProcess() error {
	sigcont := make(chan os.Signal, 1)
	signal.Notify(sigcont, syscall.SIGCONT)
	err := syscall.Kill(0, syscall.SIGTSTP)
	if err == nil {
		<-sigcont
	}
	signal.Reset(syscall.SIGCONT)
	return err
}

