package utils

import "encoding/hex"
import "crypto/md5"

func Md5(s string) string  {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

