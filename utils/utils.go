package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V 计算md5值.
func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
