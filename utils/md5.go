package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// @author: may
// @function: MD5
// @description: 密码 md5 加密
// @param: str []byte
// @return: string

func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
