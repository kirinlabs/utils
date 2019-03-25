package sys

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RealPath(f string) string {
	p, err := filepath.Abs(f)
	if err != nil {
		log.Panicln("Get absolute path error.")
	}
	p = strings.Replace(p, "\\", "/", -1)
	l := strings.LastIndex(p, "/") + 1
	return Substr(p, 0, l)
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsFile(f string) bool {
	b, err := os.Stat(f)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return false
	}
	return true
}

func IsDir(p string) bool {
	b, err := os.Stat(p)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return true
	}
	return false
}

func Substr(s string, start int, strlength ...int) string {
	charlist := []rune(s)
	l := len(charlist)
	length := 0
	end := 0

	if len(strlength) == 0 {
		length = l
	} else {
		length = strlength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])
}
