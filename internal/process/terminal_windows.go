//go:build windows

package process

// terminalForPID has no meaningful value on Windows: console processes have no
// Unix-style controlling tty, so worktree terminal identification is a no-op.
func terminalForPID(pid int32) string { return "" }
