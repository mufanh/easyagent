// +build linux darwin freebsd netbsd openbsd

package shell

func ExecuteScript(filename string, timeout int) (string, error) {
	return executeCommand("sh", []string{filename}, timeout)
}

func AsyncExecuteScript(filename string, logDir string, logFile string) error {
	return asyncExecuteCommand("sh", []string{filename}, logDir, logFile)
}

func ExecuteShell(command string, timeout int) (string, error) {
	return executeCommand("sh", []string{"-c", command}, timeout)
}

func AsyncExecuteShell(command string, logDir string, logFile string) error {
	return asyncExecuteCommand("sh", []string{"-c", command}, logDir, logFile)
}
