package utils

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
)

//Encode password
func EncodePassword(password string, salt string) string {
	md5 := md5.New()
	md5.Write([]byte(password + salt))
	md5Hex := fmt.Sprintf("%x", md5.Sum(nil))

	sha := sha512.New()
	sha.Write([]byte(md5Hex + salt))
	pwd := fmt.Sprintf("%x", sha.Sum(nil))
	return string(pwd)
}

//Md5
func Md5(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	return fmt.Sprintf("%x", md5.Sum(nil))
}
