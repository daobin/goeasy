package goeasy

import "net/http"

type context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
