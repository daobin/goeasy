package render

import "net/http"

type IRender interface {
	Render(code int, w http.ResponseWriter)
}
