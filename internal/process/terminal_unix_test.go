//go:build !windows

package process

import "testing"

func TestTerminalForPID_NoSuchProcess(t *testing.T) {
	// A PID that almost certainly does not exist must yield "", not an error.
	if got := terminalForPID(1 << 30); got != "" {
		t.Fatalf("terminalForPID(bogus) = %q, want empty", got)
	}
}

func TestTerminalForPID_NoControllingTerminal(t *testing.T) {
	// PID 1 (launchd/init) has no controlling terminal, so ps reports "?".
	if got := terminalForPID(1); got != "" {
		t.Fatalf("terminalForPID(1) = %q, want empty (no controlling tty)", got)
	}
}
