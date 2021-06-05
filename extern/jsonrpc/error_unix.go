// +build linux darwin freebsd netbsd openbsd

package jsonrpc

import "syscall"

func isConnError(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNRESET || se == syscall.ECONNABORTED
	}
	return false
}
