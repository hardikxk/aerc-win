//go:build !windows

package app

import "golang.org/x/sys/unix"

func checkAccessExecutable(path string) error {
	return unix.Access(path, unix.X_OK)
}

