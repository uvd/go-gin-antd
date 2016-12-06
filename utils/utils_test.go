package utils

import (
	"testing"
)

func TestEncodePassword(t *testing.T) {
	password := "wsc"
	salt := "123"
	salt2 := "1234"
	enpass := EncodePassword(password, salt)
	enpass2 := EncodePassword(password, salt)
	enpass3 := EncodePassword(password, salt2)
	if enpass != enpass2 {
		t.Fail()
	}

	if enpass == enpass3 {
		t.Fail()
	}
}

func TestMd5(t *testing.T) {
	src := "abc"
	Md5(src)

}
