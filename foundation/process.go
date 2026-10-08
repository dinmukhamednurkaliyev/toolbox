package foundation

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

func ExecuteProcess(commandText string, output io.Writer) error {
	var process *exec.Cmd

	if runtime.GOOS == "windows" {
		process = exec.Command("cmd", "/C", commandText)
	} else {
		process = exec.Command("sh", "-c", commandText)
	}

	process.Stdin = os.Stdin
	process.Stdout = output
	process.Stderr = os.Stderr

	if executionError := process.Run(); executionError != nil {
		return fmt.Errorf("execute command %q: %w", commandText, executionError)
	}

	return nil
}

func ProcessExitCode(executionError error) int {
	if processExitError, ok := errors.AsType[*exec.ExitError](executionError); ok {
		if processExitError.ExitCode() > 0 {
			return processExitError.ExitCode()
		}
	}

	return 1
}
