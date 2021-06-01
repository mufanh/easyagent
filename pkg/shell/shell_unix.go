// +build linux darwin freebsd netbsd openbsd

package shell

func ExecuteShell(command string, timeout int) (string, error) {
	return executeCommand("sh", []string{command}, timeout)
}

func AsyncExecuteShell(command string, logDir string, logFile string) error {
	return asyncExecuteCommand("sh", []string{command}, logDir, logFile)
}
