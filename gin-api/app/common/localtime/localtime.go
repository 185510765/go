package localtime

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// 检出 mysql 时调用  查询时使用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

// 写入 mysql 时调用 新增编辑时使用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	// fmt.Println(t.String())
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// // 时间格式化 *******************************************
// type LocalTime time.Time

// func (t *LocalTime) MarshalJSON() ([]byte, error) {
// 	tTime := time.Time(*t)
// 	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
// }

// // 查询时使用
// func (t *LocalTime) Scan(v interface{}) error {
// 	if value, ok := v.(time.Time); ok {
// 		*t = LocalTime(value)
// 		return nil
// 	}
// 	return fmt.Errorf("can not convert %v to timestamp", v)
// }

// // 新增编辑时使用
// func (t LocalTime) Value() (driver.Value, error) {
// 	var zeroTime time.Time
// 	tlt := time.Time(t)
// 	//判断给定时间是否和默认零时间的时间戳相同
// 	if tlt.UnixNano() == zeroTime.UnixNano() {
// 		return nil, nil
// 	}
// 	return tlt, nil
// }
