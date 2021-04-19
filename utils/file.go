package utils

import "os"

func IsExist(path string) bool {
	// 没有权限问题 正常
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
