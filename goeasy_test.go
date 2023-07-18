package goeasy

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("GoEasy Test >>>", func(t *testing.T) {
		got := New()
		if got == nil {
			t.Errorf("New() == %#v", got)
		}
		got.GET("/", func(c *Context) {
			c.Json(http.StatusOK, map[string]any{"name": "Index"})
		})
		got.GET("user", func(c *Context) {
			c.Json(http.StatusOK, map[string]any{"name": "User"})
		})

		fmt.Printf("GoEasy Test >>> %#v\n", got)
		Start(":8899")
	})
}
