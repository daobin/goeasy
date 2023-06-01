package goeasy

import (
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		name         string
		relativePath string
	}{
		{
			name:         "Router Handle001",
			relativePath: "/",
		},
		{
			name:         "Router Handle002",
			relativePath: "/user/list/",
		},
		{
			name:         "Router Handle002",
			relativePath: "/user",
		},
		//{
		//	name:         "Router Handle003",
		//	relativePath: "/user/:id",
		//},
		//{
		//	name:         "Router Handle004",
		//	relativePath: "/user/:id/article/list",
		//},
		//{
		//	name:         "Router Handle005",
		//	relativePath: "/user/:id/article/:id",
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New().GET(tt.relativePath, func(c *context) {})
			fmt.Printf("%s == %#v", tt.name, got)
		})
	}
}
