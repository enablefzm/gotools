package vatools

import (
	"io/ioutil"
	"strings"
)

func ReadFileLR(sFile string) ([]string, error) {
	btVal, err := ioutil.ReadFile(sFile)
	if err != nil {
		return []string{}, err
	}
	arrStr := strings.Split(string(btVal), "\r\n")
	return arrStr, nil
}

func WriteFileLR(sFile string, strVal string) error {
	return ioutil.WriteFile(sFile, []byte(strVal), 0777)
}
