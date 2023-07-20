package render

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Data any
}

func (j *Json) Render(code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	jsonBytes, _ := json.Marshal(j.Data)
	_, _ = w.Write(jsonBytes)
}

// 校验是否实现相关接口
var _ IRender = (*Json)(nil)
