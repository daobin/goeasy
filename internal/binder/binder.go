package binder

import "net/http"

type IBinder interface {
	Bind(target any, req *http.Request) error
}
