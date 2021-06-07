// +build windows

package shell

func ExecuteScript(filename string, timeout int) (string, error) {
	return executeCommand("cmd.exe", []string{"/C", filename}, timeout)
}

func AsyncExecuteScript(filename string, logDir string, logFile string) error {
	return asyncExecuteCommand("cmd.exe", []string{"/C", filename}, logDir, logFile)
}

func ExecuteShell(command string, timeout int) (string, error) {
	return executeCommand("cmd.exe", []string{"/C", command}, timeout)
}

func AsyncExecuteShell(command string, logDir string, logFile string) error {
	return asyncExecuteCommand("cmd.exe", []string{"/C", command}, logDir, logFile)
}
