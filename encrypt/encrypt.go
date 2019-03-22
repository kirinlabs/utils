package encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Md5(s string) (result string) {
	m := md5.New()
	m.Write([]byte(s))
	b := m.Sum(nil)
	result = hex.EncodeToString(b)
	return
}

func Sha1(s string) (result string) {
	sha := sha1.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

func Sha256(s string) (result string) {
	sha := sha256.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

func Sha512(s string) (result string) {
	sha := sha512.New()
	sha.Write([]byte(s))
	b := sha.Sum(nil)
	result = fmt.Sprintf("%x", b)
	return
}

func Base64Encode(s string) (result string) {
	result = base64.StdEncoding.EncodeToString([]byte(s))
	return
}

func Base64Decode(s string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "nil", err
	}
	return string(result), nil
}
