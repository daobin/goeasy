package binder

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
)

type Query struct{}

func (j *Query) Bind(target any, req *http.Request) error {
	//获取目标数据值
	targetValue := reflect.ValueOf(target)
	// 检查目标是否为非空指针
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return errors.New("请求数据解析目标必须是一个非空指针")
	}

	// 获取目标数据类型
	targetType := reflect.TypeOf(target).Elem()
	// 检查目标是否为结构体
	if targetType.Kind() != reflect.Struct {
		return errors.New("请求数据解析目标必须是结构体")
	}

	// 解析Query参数
	queryParams := req.URL.Query()

	// 遍历结构体字段并赋值
	num := targetType.NumField()
	for i := 0; i < num; i++ {
		// 获取字段标签
		tag := targetType.Field(i).Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		// 根据字段类型赋值
		fieldValue := targetValue.Elem().Field(i)
		queryValue := queryParams.Get(tag)

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(queryValue)
		case reflect.Int:
			intValue, _ := strconv.Atoi(queryValue)
			fieldValue.SetInt(int64(intValue))
		case reflect.Bool:
			boolValue, _ := strconv.ParseBool(queryValue)
			fieldValue.SetBool(boolValue)
		case reflect.Float32, reflect.Float64:
			floatValue, _ := strconv.ParseFloat(queryValue, 32)
			fieldValue.SetFloat(floatValue)
		default:
			// 其他类型赋值暂未实现
		}
	}

	return nil
}

// 校验是否实现相关接口
var _ IBinder = (*Query)(nil)
