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
		//{
		//	name:     "Router Root",
		//	fullPath: "/",
		//},
		//{
		//	name:     "Router User List",
		//	fullPath: "/user/list/",
		//},
		//{
		//	name:     "Router User Length",
		//	fullPath: "/user/length",
		//},
		{
			name:     "Router User Info",
			fullPath: "/user",
		},
		//{
		//	name:     "Router Order List",
		//	fullPath: "/order/list",
		//},
		{
			name:     "Router Handle003",
			fullPath: "/user/:id",
		},
		{
			name:     "Router Handle004",
			fullPath: "/user/:id/article/list",
		},
		{
			name:     "Router Handle005",
			fullPath: "/user/:id/article/:id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New().GET(tt.fullPath, func(c *context) {})
			fmt.Printf("%s == %#v", tt.name, got)
		})
	}
}
