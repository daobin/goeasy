package goeasy

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "New Engine",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got == nil {
				t.Errorf("New() == %#v", got)
			}
		})
	}
}
