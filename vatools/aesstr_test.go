package vatools

import (
	"fmt"
	"testing"
)

func TestAesStr(t *testing.T) {
	result := AESEnStr("fzmvava6")
	fmt.Println(result)
	resDes := AESUnStr(result)
	fmt.Println("****" + resDes + "****")

	SetPrivateAESKey("fzm1234567890abc")
	result1 := AESEnStr("fzmvava6")
	fmt.Println(result1)
	resDes1 := AESUnStr(result1)
	fmt.Println("****" + resDes1 + "****")
}
