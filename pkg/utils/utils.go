package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io/fs"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"unicode"
	"unsafe"
)

// #nosec G103
// UnsafeBytes returns a byte pointer without allocation
func UnsafeBytes(s string) []byte {
	var bs []byte

	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len

	return bs
}

// #nosec G103
// UnsafeString returns a string pointer without allocation
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsSuccess(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}

func IsInt(s string) bool {
	if len(s) > 0 && s[0] == '0' {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func Getenv(env string, def ...string) string {
	item := os.Getenv(env)

	if item == "" && len(def) > 0 {
		return def[0]
	}

	return item
}

func RandomString(n int32) string {
	buffer := make([]byte, n)

	_, err := rand.Read(buffer)
	if err != nil {
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(buffer)
}

func GetAbsolutePath(path string) (string, error) {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)

		if err != nil {
			return "", err
		}

		return path, nil
	}

	return path, err
}

func CreateDirectoryFromFile(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)
	if err != nil {
		return "", err
	}

	directory := filepath.Dir(p)

	if err := os.MkdirAll(directory, perm); err != nil {
		return "", err
	}

	return p, nil
}

func CreateFile(path string, flags int, dirMode, mode fs.FileMode) (file *os.File, err error) {
	path, err = CreateDirectoryFromFile(path, dirMode|fs.ModeDir)

	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		file, err = os.Create(path)

		if err != nil {
			return nil, err
		}

		if err := file.Chmod(mode); err != nil {
			return nil, err
		}

		if err := file.Close(); err != nil {
			return nil, err
		}
	}

	file, err = os.OpenFile(path, flags, mode)

	return
}

func CreateLogFile(path string) (file *os.File, err error) {
	file, err = CreateFile(path, os.O_WRONLY|os.O_APPEND, 0o744, fs.FileMode(0o744)|os.ModeAppend)

	return
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	return true
}

func CreateDirectory(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(p, perm); err != nil {
		return "", err
	}

	return p, nil
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)

		// check the address type and if it is not a loopback the display it
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}
