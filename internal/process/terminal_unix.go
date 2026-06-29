//go:build !windows

package process

import (
	"os/exec"
	"strconv"
	"strings"
)

// terminalForPID returns the controlling terminal of a process as a short name
// (e.g. "ttys003" on macOS, "pts/3" on Linux), or "" if it has none.
//
// We shell out to `ps` rather than use gopsutil's Process.Terminal(): that
// method is unimplemented on darwin (returns "not implemented yet"), and `ps`
// reports the tty consistently across the Unix platforms we target.
func terminalForPID(pid int32) string {
	out, err := exec.Command("ps", "-o", "tty=", "-p", strconv.Itoa(int(pid))).Output()
	if err != nil {
		return ""
	}
	tty := strings.TrimSpace(string(out))
	// ps prints "?" / "??" for a process with no controlling terminal.
	if tty == "" || tty == "?" || tty == "??" {
		return ""
	}
	return tty
}
