package vatools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

func init() {
	// 产生一个随机的种子
	rand.Seed(int64(time.Now().Nanosecond()))
}

// 获取指定的范围的随机值
func CRnd(min, max int) int {
	if max <= min {
		return min
	}
	v := max - min + 1
	res := rand.Intn(v)
	res += min
	return res
}

//	判断是否是有效的帐号ID
//  字母开头，允许5-16字节，允许字母数字下划线
//		有效返回 true
func CheckIsUID(uid string) bool {
	res, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{2,16}$", uid)
	return res
}

//	是否是有效的中文英文数字名称
//	中文、英文、数字及下划线
//		有效返回 true
func CheckIsName(name string) bool {
	res, _ := regexp.MatchString("^[\u4e00-\u9fa5a-zA-Z][\u4e00-\u9fa5_a-zA-Z0-9]{0,10}$", name)
	return res
}

// 是否是有效的密码格式
// 只能包含下划线字母和数字
//    有效返回 true
func CheckIsPass(pass string) bool {
	res, _ := regexp.MatchString("^[_a-zA-Z0-9]{3,20}$", pass)
	return res
}

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

func SUint(val string) uint {
	v := SInt(val)
	if v < 0 {
		return 0
	}
	return uint(v)
}

func SInt8(val string) int8 {
	return int8(SInt(val))
}

func SUint8(val string) uint8 {
	v := SInt(val)
	if v < 0 {
		return 0
	}
	return uint8(v)
}

func SInt16(val string) int16 {
	return int16(SInt(val))
}

// 将字符串转为UInt16
func SUint16(val string) uint16 {
	v := SInt(val)
	if v < 0 {
		return 0
	}
	return uint16(v)
}

func SInt32(val string) int32 {
	return int32(SInt(val))
}

func SUint32(val string) uint32 {
	v := SInt64(val)
	if v < 0 {
		return 0
	}
	return uint32(v)
}

func SInt64(val string) int64 {
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func SFloat64(val string) float64 {
	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return v
}

func SFloat32(val string) float32 {
	v, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0
	}
	return float32(v)
}

// 将字符串转换为整型
func SInt(val string) int {
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return v
}


// 将时间格式字符串转换为时间对象
func STime(strTime string) time.Time {
	loc, _ := time.LoadLocation("Local")
	resTime, _ := time.ParseInLocation(TIME_FORMAT, strTime, loc)
	return resTime
}

// 获取当前时间字符型格式
func GetNowTimeString() string {
	return time.Now().Format(TIME_FORMAT)
}

func GetTimeString(intTime int64) string {
	return time.Unix(intTime, 0).Format(TIME_FORMAT)
}

func MapToJson(mpInfo map[string]interface{}) string {
	btInfo, err := json.Marshal(mpInfo)
	if err != nil {
		return ""
	}
	return string(btInfo)
}

func JsonToMap(strJson string) map[string]interface{} {
	var res map[string]interface{}
	err := UnJson(strJson, &res)
	if err != nil || res == nil {
		return map[string]interface{}{}
	}
	return res
}

func Json(source interface{}) (string, error) {
	if btVal, err := json.Marshal(source); err != nil {
		return "", err
	} else {
		return string(btVal), nil
	}
}

func UnJson(strJson string, sTypeOB interface{}) error {
	return json.Unmarshal([]byte(strJson), &sTypeOB)
}
