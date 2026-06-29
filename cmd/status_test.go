package cmd

import (
	"reflect"
	"testing"

	"github.com/kunchenguid/treehouse/internal/process"
)

func TestDistinctTerminals(t *testing.T) {
	tests := []struct {
		name  string
		procs []process.ProcessInfo
		want  []string
	}{
		{
			name:  "none",
			procs: nil,
			want:  nil,
		},
		{
			name: "dedups shared terminal and drops empties",
			procs: []process.ProcessInfo{
				{Name: "zsh", Terminal: "ttys003"},
				{Name: "claude", Terminal: "ttys003"},
				{Name: "node", Terminal: ""}, // no controlling tty
			},
			want: []string{"ttys003"},
		},
		{
			name: "preserves first-seen order across terminals",
			procs: []process.ProcessInfo{
				{Name: "a", Terminal: "ttys006"},
				{Name: "b", Terminal: "ttys003"},
				{Name: "c", Terminal: "ttys006"},
			},
			want: []string{"ttys006", "ttys003"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := distinctTerminals(tc.procs); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("distinctTerminals() = %v, want %v", got, tc.want)
			}
		})
	}
}
