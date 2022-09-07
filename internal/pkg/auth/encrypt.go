package auth

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

const saltLength = 4

func MakeSalt() string {
	now := time.Now().Nanosecond() - time.Now().Second()
	nowBuf := []byte(strconv.FormatInt(int64(now), 10))

	hash := md5.New()
	hash.Write(nowBuf)

	salt := fmt.Sprintf("%x", hash.Sum(nil))
	salt = salt[:saltLength]
	return salt
}

func Encrypt(password, salt string) string {
	buf := []byte(password + salt)
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
