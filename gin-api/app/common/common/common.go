package common

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
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

// 金额补0 5.00 5.50
func PriceFillZero(price string) string {
	priceLen := len(price)
	if priceLen < 4 {
		if priceLen == 1 {
			return price + ".00"
		}
		if priceLen == 3 {
			return price + "0"
		}
	}

	return price
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

// 获取昨天开始时间戳和结束时间戳
func GetYesterdayStartTimeAndEndTime() (int64, int64) {
	NowTime := time.Now()
	var startTime time.Time
	//NowTime := time.Date(2022, 9, 15, 0, 0, 0, 0, time.Local)
	if NowTime.Hour() == 0 && NowTime.Minute() == 0 && NowTime.Second() == 0 {
		startTime = time.Unix(NowTime.Unix()-86399, 0) //当天的最后一秒
	} else {
		startTime = time.Unix(NowTime.Unix()-86400, 0)
	}
	currentYear := startTime.Year()
	currentMonth := startTime.Month()
	currentDay := startTime.Day()
	yesterdayStartTime := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.Local).Unix()
	yesterdayEndTime := time.Date(currentYear, currentMonth, currentDay, 23, 59, 59, 0, time.Local).Unix()
	//fmt.Println(yesterdayStartTime, yesterdayEndTime)
	return yesterdayStartTime, yesterdayEndTime
}

// 获取今天开始时间戳和结束时间戳
func GetTodayStartTimeAndEndTime() (int64, int64) {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()
	todayStartTime := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.Local).Unix()
	todayEndTime := time.Date(currentYear, currentMonth, currentDay, 23, 59, 59, 0, time.Local).Unix()
	return todayStartTime, todayEndTime
}

// GetIp 获取本地IP地址 利用udp
func GetIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.1.20:61085
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip
}

// 地址转整形
func IpToInt(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])

}

// 获取Mac地址
func GetMacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	if len(interfaces) == 0 {
		return ""
	}

	maxIndexInterface := interfaces[0]
	for _, inter := range interfaces {
		if inter.HardwareAddr == nil {
			continue
		}
		if inter.Flags&net.FlagUp == 1 {
			maxIndexInterface = inter
		}
	}
	return maxIndexInterface.HardwareAddr.String()
}

// 发送邮件
func SendMail(servername, fromStr, password, toStr, subject, body, mailtype string) error {
	from := mail.Address{Name: "", Address: fromStr}
	to := mail.Address{Name: "", Address: toStr}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += contentType + "\r\n" + body

	// message_bak := "From: " + from.String() + "\r\n" + "To: " + to.String() + "\r\n" + "Subject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body + "\r\n"

	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", from.Address, password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}
	d, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	// Auth
	if err = d.Auth(auth); err != nil {
		return err
	}
	// To && From
	if err = d.Mail(from.Address); err != nil {
		return err
	}
	if err = d.Rcpt(to.Address); err != nil {
		return err
	}
	// Data
	w, err := d.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return d.Quit()
}
