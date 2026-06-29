package process

import "testing"

func TestProcessInfoString(t *testing.T) {
	tests := []struct {
		name string
		info ProcessInfo
		want string
	}{
		{
			name: "with terminal",
			info: ProcessInfo{PID: 123, Name: "zsh", Terminal: "ttys003"},
			want: "zsh (123, ttys003)",
		},
		{
			name: "without terminal",
			info: ProcessInfo{PID: 456, Name: "node"},
			want: "node (456)",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.info.String(); got != tc.want {
				t.Fatalf("String() = %q, want %q", got, tc.want)
			}
		})
	}
}
