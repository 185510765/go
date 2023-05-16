package common

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogf/gf/util/gconv"
)

// 获取本周一时间戳
func GetFirstDateOfWeekTS() (ts int64) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	ts = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()
	return
}

// 获取周几时间戳
func GetWeekTS(day int64) (ts int64) {
	thisWeekMonday := GetFirstDateOfWeekTS()
	ts = thisWeekMonday + (day)*86400
	return
}

// 字符串长度
func StrLen(str string) int {
	return utf8.RuneCountInString(str)
}

// 截取字符串
func StrSub(str string, sub ...int) string {
	start := sub[0]
	length := 0
	if len(sub) > 1 {
		length = sub[1]
	}

	if length < 1 {
		return string(([]rune(str))[start:])
	}
	return string(([]rune(str))[start:length])
}

// 合并字符串
func StrCombine(str ...string) string {
	var bt bytes.Buffer
	for _, arg := range str {
		bt.WriteString(arg)
	}
	//获得拼接后的字符串
	return bt.String()
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func ToUnicode(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	return textQuoted[1 : len(textQuoted)-1]
}

func UnicodeTo(str string) string {
	sUnicodev := strings.Split(str, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}

// 获取当天0点时间抽
func TodayTS() int64 {
	now := time.Now()
	return GetZeroTime(now).Unix()
}

func TodayDate() string {
	return time.Now().Format("2006/01/02")
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func HttpGet(iUrl string, params map[string]string) ([]byte, error) {
	query := ""
	for k, v := range params {
		query = fmt.Sprintf("%s%s=%s&", query, k, url.QueryEscape(v))
	}

	client := http.Client{Timeout: 5 * time.Second} //创建客户端
	resp, err := client.Get(fmt.Sprintf("%s?%s", iUrl, query))
	fmt.Println(iUrl, query)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

// 类似php http_build_query
func HttpBuildQuery(params map[string]interface{}, parentKey string) (param_str string) {
	//fmt.Println("parentKey ", parentKey)
	params_arr := make([]string, 0)
	for k, v := range params {
		if vals, ok := v.(map[string]interface{}); ok {
			if parentKey != "" {
				k = fmt.Sprintf("%s[%s]", parentKey, k)
			}
			params_arr = append(params_arr, HttpBuildQuery(vals, k))
		} else {
			if parentKey != "" {
				params_arr = append(params_arr, fmt.Sprintf("%s[%s]=%s", parentKey, k, gconv.String(v)))
			} else {
				params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, gconv.String(v)))
			}
		}
	}
	param_str = strings.Join(params_arr, "&")
	return param_str
}

// 随机ip
func GenIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

// interface 转 string
func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// 去除价格中的中文、斜杠/、钱符号
func GetPrice(price string) string {
	value := "0.00"

	if price != "" {
		reg := regexp.MustCompile("[\u4e00-\u9fa5]") // 中文
		result := reg.ReplaceAllString(price, "")
		// fmt.Println(result)

		reg1 := regexp.MustCompile("/")
		result1 := reg1.ReplaceAllString(result, "")
		// fmt.Println(result1)

		reg2 := regexp.MustCompile("¥")
		result2 := reg2.ReplaceAllString(result1, "")
		// fmt.Println(result2)

		value = result2
	}

	return value
}

// 比较两个map是否相等
func CompareTwoMapInterface(data1 map[string]interface{}, data2 map[string]interface{}) bool {
	keySlice := make([]string, 0)
	dataSlice1 := make([]interface{}, 0)
	dataSlice2 := make([]interface{}, 0)

	for key, value := range data1 {
		keySlice = append(keySlice, key)
		dataSlice1 = append(dataSlice1, value)
	}
	for _, key := range keySlice {
		if data, ok := data2[key]; ok {
			dataSlice2 = append(dataSlice2, data)
		} else {
			return false
		}
	}

	dataStr1, _ := json.Marshal(dataSlice1)
	dataStr2, _ := json.Marshal(dataSlice2)

	return string(dataStr1) == string(dataStr2)
}
