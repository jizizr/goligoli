package tools

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string, salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str+salt)))
}
