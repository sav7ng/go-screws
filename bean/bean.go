package bean

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"runtime/debug"
	"strings"
)

type BeanTool struct{}

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

func ToB[S, T any](s *S) *T {
	if s == nil {
		return nil
	}
	var t T
	data, err := jsonIter.Marshal(*s)
	if err != nil {
		debug.PrintStack()
		return &t
	}
	err = jsonIter.Unmarshal(data, &t)
	if err != nil {
		debug.PrintStack()
	}
	return &t
}

func ToBs[S, T any](s *[]S) *[]T {
	if s == nil || len(*s) < 1 {
		return nil
	}
	var t []T
	data, err := jsonIter.Marshal(*s)
	if err != nil {
		debug.PrintStack()
		return &t
	}
	err = jsonIter.Unmarshal(data, &t)
	if err != nil {
		debug.PrintStack()
	}
	return &t
}

func ToBs2[S, T any](s *[]*S) *[]*T {
	if s == nil || len(*s) < 1 {
		return nil
	}
	var t []*T
	data, err := jsonIter.Marshal(*s)
	if err != nil {
		debug.PrintStack()
		return &t
	}
	err = jsonIter.Unmarshal(data, &t)
	if err != nil {
		debug.PrintStack()
	}
	return &t
}

func ToBs3[S, T any](s *[]S) *[]*T {
	if s == nil || len(*s) < 1 {
		return nil
	}
	var t []*T
	data, err := jsonIter.Marshal(*s)
	if err != nil {
		debug.PrintStack()
		return &t
	}
	err = jsonIter.Unmarshal(data, &t)
	if err != nil {
		debug.PrintStack()
	}
	return &t
}

// ToBs4 参数为空时不返回nil而是返回空切片
func ToBs4[S, T any](s *[]S) *[]*T {
	if s == nil || len(*s) < 1 {
		result := make([]*T, 0)
		return &result
	}
	return ToBs3[S, T](s)
}

// ToMap 利用json中转的方式将结构体转map
func ToMap[S any](s *S) *map[string]interface{} {
	jsonBytes, err := jsonIter.Marshal(*s)
	if err != nil {
		return nil
	}
	var data map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(string(jsonBytes)))
	decoder.UseNumber() // 通过这个方法指定使用 json.Number 类型来表示数字
	err = decoder.Decode(&data)
	if err != nil {
		return nil
	}
	return &data
	// 用这种方式转map的话会把int64当做float64去处理，导致精度丢失
	//return ToB[S, map[string]interface{}](s)
}

func ToJs[S any](s *S) *string {
	if s == nil {
		return nil
	}
	data, err := jsonIter.MarshalToString(*s)
	if err != nil {

		debug.PrintStack()
	}
	return &data
}

func JsToB[T any](s *string) *T {
	if s == nil {
		return nil
	}
	var t T
	err := jsonIter.UnmarshalFromString(*s, &t)
	if err != nil {

		debug.PrintStack()
	}
	return &t
}

func JsToBs[T any](s *string) *[]T {
	if s == nil {
		return nil
	}
	var t []T
	err := jsonIter.UnmarshalFromString(*s, &t)
	if err != nil {

		debug.PrintStack()
	}
	return &t
}
