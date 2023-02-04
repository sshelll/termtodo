package util

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetKeyRune(key string) string {
	if !strings.HasPrefix(key, "Rune") {
		return key
	}
	return strings.ReplaceAll(strings.ReplaceAll(key, "]", ""), "Rune[", "")
}

func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func JoinHomePath(path string) (string, error) {

	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(u.HomeDir, path), nil

}

func WithFatalf(fn func() error, funcName string) {
	if err := fn(); err != nil {
		log.Fatalf("exec %s failed: %v", funcName, err)
	}
}

func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func ContainStr(src []string, target string) bool {
	for _, s := range src {
		if target == s {
			return true
		}
	}
	return false
}

func CreateFile(path string) error {
	if Stat(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

func Stat(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
