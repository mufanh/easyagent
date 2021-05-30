// +build windows

package shell

func ExecuteShell(command string, timeout int) (string, error) {
	return executeCommand("cmd.exe", []string{"/C", command}, timeout)
}

func AsyncExecuteShell(command string, logDir string, logFile string) error {
	return asyncExecuteCommand("cmd.exe", []string{"/C", command}, logDir, logFile)
}
