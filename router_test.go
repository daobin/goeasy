package goeasy

import (
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		name     string
		fullPath string
	}{
		{
			name:     "/admin/user/info  [yes]",
			fullPath: "/admin/user/info",
		},
		{
			name:     "/admin/user/:id  [yes]",
			fullPath: "/admin/user/:id",
		},
		{
			name:     "/admin/user/:name  [no]",
			fullPath: "/admin/user/:name",
		},
		{
			name:     "/admin/:object/:id  [yes]",
			fullPath: "/admin/:object/:id",
		},
		{
			name:     "/admin/:object/*key  [no]",
			fullPath: "/admin/:object/*key",
		},
		{
			name:     "/admin/:object/list  [no]",
			fullPath: "/admin/:object/list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New().GET(tt.fullPath, func(c *Context) {})
			fmt.Printf("%s == %#v", tt.name, got)
		})
	}
}
