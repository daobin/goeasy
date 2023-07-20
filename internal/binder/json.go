package binder

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Json struct{}

func (j *Json) Bind(target any, req *http.Request) error {
	// 校验请求数据
	if req == nil || req.Body == nil {
		return errors.New("请求数据无效")
	}

	// 获取请求数据
	reqBody, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	// 解析请求数据
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(target); err != nil {
		log.Println("请求数据解析失败：", err.Error())
		return errors.New("请求数据解析失败")
	}

	// 重置请求数据，以便可以重复调用该方法获取请求数据
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	return nil
}

// 校验是否实现相关接口
var _ IBinder = (*Json)(nil)
