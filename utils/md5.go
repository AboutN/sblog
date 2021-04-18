package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(origData string) string {
	hash := md5.New()
	hash.Write([]byte(origData))
	return hex.EncodeToString(hash.Sum(nil))
}
