package scode

import "strconv"

// StatusCode .
type StatusCode int

const (
	// Success 成功
	Success StatusCode = 200000
	// Failure 失败
	Failure StatusCode = 400000
)

// 指向不是太具体额错误描述
const (
	// ErrorRequire 必须项
	ErrorDB StatusCode = iota + 400001
)

func (sc StatusCode) String() string {
	switch sc {
	case Success:
		return "成功"
	case Failure:
		return "失败"
	case ErrorDB:
		return "数据库错误"
	}
	return "未知: " + strconv.FormatInt(int64(sc), 10)
}
