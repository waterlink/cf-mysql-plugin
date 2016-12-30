package interactivity

import (
	"fmt"
	"os"
)

// Detector is for determining if STDIN/STDOUT is in interactive or pipe mode
type Detector interface {
	IsInteractive() bool
}

// FileStatDetector is for determining if STDIN/STDOUT is in interactive or pipe
// mode using stat() syscall
type FileStatDetector struct{}

// IsInteractive returns true if STDIN or STDOUT is interactive.
// Otherwise returns false.
func (FileStatDetector) IsInteractive() bool {
	return isInteractive(os.Stdout) && isInteractive(os.Stdin)
}

func isInteractive(file *os.File) bool {
	info, err := file.Stat()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to detect if STDIN is interactive. Assuming non-interactive. Error: %s", err)
		return false
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		return false
	}

	return true
}
