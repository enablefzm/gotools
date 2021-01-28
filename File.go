package gotools

import (
	"io/ioutil"
	"strings"
)

func ReadFileLR(sFile string) ([]string, error) {
	btVal, err := ioutil.ReadFile(sFile)
	if err != nil {
		return nil, err
	}
	arrStr := strings.Split(string(btVal), "\r\n")
	return arrStr, nil
}
