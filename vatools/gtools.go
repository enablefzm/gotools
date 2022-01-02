package vatools

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
)

func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

// 获取当前运的目录
func GetNowPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "/", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "/", err
	}
	return filepath.Dir(path), nil
}
